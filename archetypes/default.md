---
title: "{{ replace .TranslationBaseName "-" " " | title }}"
author: 邹慧刚
date: {{ dateFormat "2006-01-02 15:04:05" .Date }}
categories:
  - "{{ replace .Dir "\\" "" }}"
tags: ["post", "{{ replace .TranslationBaseName "-" " " | title }}"]
draft: false
---

