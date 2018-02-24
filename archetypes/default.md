---
title: "{{ replace .TranslationBaseName "-" " " | title }}"
description: "这是一段描述-zouhuigang.anooc.com"
author: 邹慧刚
date: {{ .Date }}
categories:
  - "{{ replace .Dir "\\" "" }}"
tags: ["post", "{{ replace .TranslationBaseName "-" " " | title }}"]
draft: false
---