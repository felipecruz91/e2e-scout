# e2e-scout

This is a demo project that consists of a Go binary with some vulnerable packages. The binary is packaged as a Docker image whose base image is a pinned version of an alpine image with OS-specific vulnerabilities.


## Build the image

```bash
docker buildx build -t e2e-scout . --load
```
The image size is 24.8MB.

## Analyze the image

```
docker scout cves e2e-scout
    ✓ Image stored for indexing
    ✓ Indexed 36 packages
    ✗ Detected 4 vulnerable packages with a total of 4 vulnerabilities


## Overview

                    │       Analyzed Image
────────────────────┼──────────────────────────────
  Target            │  e2e-scout:latest
    digest          │  09248ddbe821
    platform        │ linux/arm64/v8
    vulnerabilities │    1C     1H     1M     1L
    size            │ 6.7 MB
    packages        │ 36


## Packages and Vulnerabilities

   1C     0H     0M     0L  openssl 3.3.1-r0
pkg:apk/alpine/openssl@3.3.1-r0?os_name=alpine&os_version=3.20

    ✗ CRITICAL CVE-2024-5535
      https://scout.docker.com/v/CVE-2024-5535
      Affected range  : <3.3.1-r1
      Fixed version   : 3.3.1-r1
      EPSS Score      : 0.04%
      EPSS Percentile : 13th percentile


   0C     1H     0M     0L  stdlib 1.22.4
pkg:golang/stdlib@1.22.4

    ✗ HIGH CVE-2024-24791
      https://scout.docker.com/v/CVE-2024-24791
      Affected range  : >=1.22.0-0
                      : <1.22.5
      Fixed version   : 1.22.5
      EPSS Score      : 0.04%
      EPSS Percentile : 16th percentile


   0C     0H     1M     0L  github.com/hashicorp/go-retryablehttp 0.7.6
pkg:golang/github.com/hashicorp/go-retryablehttp@0.7.6

    ✗ MEDIUM CVE-2024-6104 [Insertion of Sensitive Information into Log File]
      https://scout.docker.com/v/CVE-2024-6104
      Affected range  : <0.7.7
      Fixed version   : 0.7.7
      CVSS Score      : 6.0
      CVSS Vector     : CVSS:3.1/AV:L/AC:L/PR:H/UI:N/S:C/C:H/I:N/A:N
      EPSS Score      : 0.04%
      EPSS Percentile : 9th percentile


   0C     0H     0M     1L  github.com/tektoncd/pipeline 0.52.0
pkg:golang/github.com/tektoncd/pipeline@0.52.0

    ✗ LOW CVE-2023-37264 [Insufficient Verification of Data Authenticity]
      https://scout.docker.com/v/CVE-2023-37264
      Affected range  : >=0.35.0
                      : <=0.52.0
      Fixed version   : not fixed
      CVSS Score      : 3.7
      CVSS Vector     : CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:L/I:L/A:N
      EPSS Score      : 0.07%
      EPSS Percentile : 31st percentile



4 vulnerabilities found in 4 packages
  LOW       1
  MEDIUM    1
  HIGH      1
  CRITICAL  1
```
