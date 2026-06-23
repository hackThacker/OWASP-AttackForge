<?php
// Disable strict exception reporting for mysqli (required by older vulnerable targets)
if (function_exists('mysqli_report')) {
    mysqli_report(MYSQLI_REPORT_OFF);
}
?>
