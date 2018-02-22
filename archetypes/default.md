---
title: "{{ replace .TranslationBaseName "-" " " | title }}"
author: zouhuigang {{.LogicalName}} {{.Path}} {{.Dir}}
date: {{ dateFormat "2006-01-02 15:04:05" .Date }}
banner: img/indiagate1.jpg
categories:
  - "{{ replace .Dir "\\" "" }}"
tags: ["post", "{{ replace .TranslationBaseName "-" " " | title }}"]
draft: false
---

