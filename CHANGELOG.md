# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2026-06-25


### Bug Fixes

- *(build)* Correct juiceshop, zerohealth-server, and restaurant Dockerfile build issues ([476829c](https://github.com/hackThacker/OWASP-AttackForge/commit/476829cfc63b48e0077c68559f053f227fb13c66))

- *(build)* Remove BuildKit cache mount and correct sbom-action input ([6ce51e2](https://github.com/hackThacker/OWASP-AttackForge/commit/6ce51e214c89566498f4c67d6202c9e0293802fb))

- *(attest)* Align service tags and retrieve image digests for SLSA attestation ([708484c](https://github.com/hackThacker/OWASP-AttackForge/commit/708484c3361e10649e51bcfd014f26044468b66b))


### Miscellaneous Chores

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([6bfbd73](https://github.com/hackThacker/OWASP-AttackForge/commit/6bfbd7378948c1aa027ad74780a8f17157d3b3db))

- *(changelog)* Update CHANGELOG.md for v1.2.0 [skip ci] ([b8c1429](https://github.com/hackThacker/OWASP-AttackForge/commit/b8c1429531a708648dcb5a5c8ee171d8795a63b1))

- *(changelog)* Update CHANGELOG.md for v1.2.0 [skip ci] ([c7dbe53](https://github.com/hackThacker/OWASP-AttackForge/commit/c7dbe5372711bce774123ee6ba7d912270700697))


## [1.1.0] - 2026-06-25


### Bug Fixes

- Fix Trivy action tag reference version prefix ([8076ff1](https://github.com/hackThacker/OWASP-AttackForge/commit/8076ff18e8e2bcf1fe415057457aadb188b9a561))

- Update Trivy action version and configure security workflows as audit-only to prevent blocking releases for intentionally vulnerable lab environments ([45b6939](https://github.com/hackThacker/OWASP-AttackForge/commit/45b6939cd625744ea933c946b89d34e932d0bf34))

- Correct conditional syntax in security.yml ([54687e2](https://github.com/hackThacker/OWASP-AttackForge/commit/54687e2088185fd2810caad5a7819536b8545876))

- Resolve conditional syntax error in security.yml with continue-on-error ([a2845ec](https://github.com/hackThacker/OWASP-AttackForge/commit/a2845ec9a3d9396bc337053aa02e86617b3e00fb))

- Ignore Docker Scout entitlement errors during scanning ([32af02a](https://github.com/hackThacker/OWASP-AttackForge/commit/32af02a26eba17e045b219e5959588d0dc89589b))

- Resolve keycloak startup race, localhost IPv6 mismatch, restaurant db compilation, and endpoint checks ([9835fee](https://github.com/hackThacker/OWASP-AttackForge/commit/9835fee931e7f649b14d804e5e68b6f9d72f01c4))

- *(ci)* Remove top-level build concurrency, fix changelog generation template, and update service/container counts in README ([775e329](https://github.com/hackThacker/OWASP-AttackForge/commit/775e32920bfb009ad1872036f20fd5ad38b3c542))

- *(ci)* Untrack installer files and add them to .gitignore ([6f0ee3d](https://github.com/hackThacker/OWASP-AttackForge/commit/6f0ee3d03990a09d9ad2136262e5d1dc0f08de29))

- *(ci)* Define explicit job permissions in release workflow and fix detached HEAD push in changelog workflow ([d825b90](https://github.com/hackThacker/OWASP-AttackForge/commit/d825b90b03653e1bfdf367edbe7d54cf6d7eccf3))

- *(ci)* Support conditional Docker Hub publishing and specify RELEASE_NOTES.md as body_path for github releases ([2e3cbc1](https://github.com/hackThacker/OWASP-AttackForge/commit/2e3cbc16807b6cab1d88edf15b2a173c9e3049f5))

- *(ci)* Correct conditional if expression syntax in build workflow ([1fdcbb9](https://github.com/hackThacker/OWASP-AttackForge/commit/1fdcbb964219a44dbf93296d6f6b45419b83819e))

- *(ci)* Declare top-level permissions on release workflow and remove job-level permissions ([6674ff9](https://github.com/hackThacker/OWASP-AttackForge/commit/6674ff96c2caa52f6b0332ed5761c0e5975c176b))

- *(ci)* Use shell-based conditional login check for Docker Hub in build.yml ([d291d24](https://github.com/hackThacker/OWASP-AttackForge/commit/d291d24d5247649343c1e1b54af77a185fb28645))

- *(release)* Use artifact sharing to carry notes draft to verification job ([cb85260](https://github.com/hackThacker/OWASP-AttackForge/commit/cb85260286a600bc38e811eb09782b551dab3bfe))

- *(release)* Enforce lowercase registry namespace owner for tags and attestations ([b5f3bc2](https://github.com/hackThacker/OWASP-AttackForge/commit/b5f3bc27c2ba0efee5231ede72728c4c4d86c608))


### Documentation

- Update README.md to document 18-app stack, local domains, credentials, and folder structures ([114f138](https://github.com/hackThacker/OWASP-AttackForge/commit/114f138e4e3d0382cb33c99228f4cf3518c46293))

- *(labs)* Add comprehensive application metadata ([fc81922](https://github.com/hackThacker/OWASP-AttackForge/commit/fc81922dc9e44dc761b42fa94e9d81253b51e167))


### Features

- Implement Enterprise DevSecOps CI/CD Platform ([1b05697](https://github.com/hackThacker/OWASP-AttackForge/commit/1b056975d488e75c63cd82ae72bef27dbd01807f))

- Integrate seven new vulnerable applications to the cyber range ([b25ed9c](https://github.com/hackThacker/OWASP-AttackForge/commit/b25ed9cc85f3279168722e8b5da9524c3120fa8a))

- *(release)* Refactor release management and automated package verification pipeline ([c466caf](https://github.com/hackThacker/OWASP-AttackForge/commit/c466caf3545e2b9446dffd35febb44910646a084))


### Miscellaneous Chores

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([29318f7](https://github.com/hackThacker/OWASP-AttackForge/commit/29318f71607c047230515e1d9be041ad63af1ef4))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([912a1db](https://github.com/hackThacker/OWASP-AttackForge/commit/912a1db2aa778a45a0dc3f0fd056fe0a3d34afb5))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([b9d6122](https://github.com/hackThacker/OWASP-AttackForge/commit/b9d612250dd7dba72fd63bf305443e76930597d6))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([a0912bd](https://github.com/hackThacker/OWASP-AttackForge/commit/a0912bdb6ea3e453410e64f74c15a66eafe42bf8))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([bc8ddc7](https://github.com/hackThacker/OWASP-AttackForge/commit/bc8ddc71482a80c8b1145dfc9666e236c1fd02aa))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([ca7f095](https://github.com/hackThacker/OWASP-AttackForge/commit/ca7f0959540246c4c0589541033eeeb84ab42dd8))

- *(changelog)* Update CHANGELOG.md for v1.1.0 [skip ci] ([b553f8f](https://github.com/hackThacker/OWASP-AttackForge/commit/b553f8fb42c98511506d58ce2f0b28212c2793f4))


## [1.0.0] - 2026-06-24


### Bug Fixes

- Resolve OWASP lab database, PHP config, and container permission issues ([687e43d](https://github.com/hackThacker/OWASP-AttackForge/commit/687e43d1911c3da13416068e2ee376bc7d29816d))


### Features

- *(docker)* Add multi-container OWASP lab deployment stack ([32b0191](https://github.com/hackThacker/OWASP-AttackForge/commit/32b019170c0587729ed99400637f653468e57c61))

- *(deployment)* Improve stability on low-memory Linux hosts ([d915b79](https://github.com/hackThacker/OWASP-AttackForge/commit/d915b792c9beb9d9c9684c7c46da37cab004f638))

- Implement production-ready CI/CD pipelines, integrate security shepherd, wrongsecrets, and rebrand to hackthacker-labs ([89961a2](https://github.com/hackThacker/OWASP-AttackForge/commit/89961a2fc97f1caaf52f5f31fb392816381df986))

<!-- generated by git-cliff -->
