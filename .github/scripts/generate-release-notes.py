import subprocess
import os
import sys
import re

def run_git(args):
    try:
        res = subprocess.run(["git"] + args, capture_output=True, text=True, check=True)
        return res.stdout.strip()
    except subprocess.CalledProcessError as e:
        print(f"Error running git {' '.join(args)}: {e.stderr}", file=sys.stderr)
        return ""

def get_release_name(tag_name):
    # Try to get annotated tag subject
    subject = run_git(["tag", "-l", "--format=%(contents:subject)", tag_name])
    if subject and not subject.startswith("v"):
        return subject
    
    # If tag has a message but it matches tag name, look for a second line/body
    body = run_git(["tag", "-l", "--format=%(contents:body)", tag_name])
    if body:
        first_line = body.strip().split("\n")[0].strip()
        if first_line:
            return first_line
            
    # Default fallback
    return "Stable Container Release"

def main():
    tag_name = os.environ.get("GITHUB_REF_NAME")
    if not tag_name:
        # Fallback for local testing
        tag_name = run_git(["describe", "--tags", "--abbrev=0"])
        if not tag_name:
            print("Error: GITHUB_REF_NAME not set and no tags found.", file=sys.stderr)
            sys.exit(1)
            
    print(f"Generating release notes for tag: {tag_name}")
    
    # Find previous tag
    previous_tag = run_git(["describe", "--tags", "--abbrev=0", f"{tag_name}^"])
    
    if previous_tag:
        print(f"Comparing {previous_tag} with {tag_name}")
        commit_range = f"{previous_tag}..{tag_name}"
    else:
        print("No previous tag found. Getting all commits.")
        commit_range = tag_name

    # Get commits formatted
    git_log_output = run_git(["log", "--format=%H|%s", commit_range])
    commits = []
    if git_log_output:
        for line in git_log_output.split("\n"):
            if "|" in line:
                sha, msg = line.split("|", 1)
                commits.append((sha, msg))

    # Categories list
    added = []
    improved = []
    fixed = []
    security = []
    infra = []

    # Conventional commits regex
    # Matches prefixes like feat(scope):, feat:, fix:, etc.
    pattern = re.compile(r"^([a-z0-9_]+)(?:\([a-z0-9_-]+\))?!?\s*:\s*(.*)$", re.IGNORECASE)

    for sha, msg in commits:
        match = pattern.match(msg)
        if match:
            category_prefix = match.group(1).lower()
            clean_msg = match.group(2).strip()
            # Capitalize first letter
            clean_msg = clean_msg[0].upper() + clean_msg[1:] if clean_msg else ""
            item = f"- {clean_msg} ([{sha[:7]}](https://github.com/hackThacker/OWASP-AttackForge/commit/{sha}))"
            
            if category_prefix == "feat":
                added.append(item)
            elif category_prefix in ("refactor", "perf", "docs", "style"):
                improved.append(item)
            elif category_prefix == "fix":
                fixed.append(item)
            elif category_prefix == "security":
                security.append(item)
            elif category_prefix in ("ci", "build"):
                infra.append(item)
        else:
            # Fallback for unconventional commits
            item = f"- {msg.strip()} ([{sha[:7]}](https://github.com/hackThacker/OWASP-AttackForge/commit/{sha}))"
            # Attempt to classify based on keywords
            msg_lower = msg.lower()
            if any(k in msg_lower for k in ("add", "new", "introduce", "implement")):
                added.append(item)
            elif any(k in msg_lower for k in ("fix", "bug", "resolve", "patch", "error")):
                fixed.append(item)
            elif any(k in msg_lower for k in ("security", "vuln", "cve", "auth", "token", "leak")):
                security.append(item)
            elif any(k in msg_lower for k in ("ci", "build", "docker", "workflow", "compose")):
                infra.append(item)
            else:
                improved.append(item)

    # Release Name
    release_name = get_release_name(tag_name)
    title = f"{tag_name} – {release_name}"
    print(f"Formatted Title: {title}")

    # Generate Markdown
    md = []
    md.append(f"# {title}\n")
    md.append("## Overview")
    md.append("This release expands the AttackForge training platform with additional vulnerable applications, improved deployment automation, and infrastructure stability enhancements.\n")

    if added:
        md.append("## Added")
        md.extend(added)
        md.append("")
        
    if improved:
        md.append("## Improved")
        md.extend(improved)
        md.append("")
        
    if fixed:
        md.append("## Fixed")
        md.extend(fixed)
        md.append("")
        
    if security:
        md.append("## Security")
        md.extend(security)
        md.append("")
        
    if infra:
        md.append("## Infrastructure")
        md.extend(infra)
        md.append("")

    md.append("## Upgrade Notes")
    md.append("Pull the latest container images and redeploy using docker-compose:\n")
    md.append("```bash")
    md.append("docker compose pull")
    md.append("docker compose up -d")
    md.append("```")

    notes_content = "\n".join(md)
    
    # Save to file
    with open("RELEASE_NOTES_DRAFT.md", "w", encoding="utf-8") as f:
        f.write(notes_content)
        
    # Output properties for GitHub Actions
    if "GITHUB_OUTPUT" in os.environ:
        with open(os.environ["GITHUB_OUTPUT"], "a") as gh_out:
            gh_out.write(f"release_title={title}\n")
            
    print("Successfully generated RELEASE_NOTES_DRAFT.md")

if __name__ == "__main__":
    main()
