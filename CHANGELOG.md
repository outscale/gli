# 📜 Changelog

## [v0.0.12] - 2026-02-24

### ✨ Added
* ✨ feat(iaas): add aliases (link/unlink, ...) by @jfbus in [#74](https://github.com/outscale/octl/pull/74)

### 🛠️ Changed / Refactoring
* 💄 ui: markdown command help + shorter flag help by @jfbus in [#75](https://github.com/outscale/octl/pull/75)

## [v0.0.11] - 2026-02-23

### ✨ Added
* ✨ feat(flags): add file based flags (keys, policies) by @jfbus in [#72](https://github.com/outscale/octl/pull/72)

### 🐛 Fixed
* 🥅 errors: hide github api errors by @jfbus in [#70](https://github.com/outscale/octl/pull/70)
* 🐛 fix(profile): fix profile loading priority by @jfbus in [#71](https://github.com/outscale/octl/pull/71)

## [v0.0.10] - 2026-02-20

### ✨ Added
* ✨ feat(iaas): iterate over API pages by @jfbus in [#57](https://github.com/outscale/octl/pull/57)
* ✨ feat: add filters, move jq to filter by @jfbus in [#60](https://github.com/outscale/octl/pull/60)
* ✨ feat(iaas): allow multiple describe/updates/deletes by @jfbus in [#62](https://github.com/outscale/octl/pull/62)
* ✨ feat(iaas): vm start/stop by @jfbus in [#68](https://github.com/outscale/octl/pull/68)
* ✨ feat(iaas): vm readconsole by @jfbus in [#69](https://github.com/outscale/octl/pull/69)

### 🛠️ Changed / Refactoring
* 🚸 ux(profile): autocomplete --profile with profile names by @jfbus in [#58](https://github.com/outscale/octl/pull/58)
* 🚸 ux(catalog): add columns for catalog + sort table by @jfbus in [#63](https://github.com/outscale/octl/pull/63)

### 🐛 Fixed
* 🐛 fix: single responses were converted to lists by @jfbus in [#61](https://github.com/outscale/octl/pull/61)
* 💩 poop: hardcode a pagination limit to avoid infinite loops by @jfbus in [#67](https://github.com/outscale/octl/pull/67)

## [v0.0.9] - 2026-02-18

### ✨ Added
* ✨ feat(iaas): add create/update aliases by @jfbus in [#51](https://github.com/outscale/octl/pull/51)
* ✨ feat(iaas): more table definitions (apilogs, ...) by @jfbus in [#54](https://github.com/outscale/octl/pull/54)
* ✨ feat(iaas): quotas & exploded lists by @jfbus in [#56](https://github.com/outscale/octl/pull/56)

### 🐛 Fixed
* 🐛 fix(spinner): cursor not reappearing after spinner display by @jfbus in [#52](https://github.com/outscale/octl/pull/52)
* 🐛 fix(iaas): some flags were wrongly marked as required by @jfbus in [#53](https://github.com/outscale/octl/pull/53)

### 📦 Dependency updates
* ⬆️ deps(gomod): update module github.com/expr-lang/expr to v1.17.8 by @Open-Source-Bot in [#55](https://github.com/outscale/octl/pull/55)
* ⬆️ deps(gomod): update module golang.org/x/mod to v0.33.0 by @Open-Source-Bot in [#40](https://github.com/outscale/octl/pull/40)

## [v0.0.8] - 2026-02-17

### ✨ Added
* ✨ feat(columns): + to add columns to output by @jfbus in [#44](https://github.com/outscale/octl/pull/44)
* ✨ feat: add profile management commands by @jfbus in [#45](https://github.com/outscale/octl/pull/45)

### 🛠️ Changed / Refactoring
* 🚸 ux: add spinner on long API calls by @jfbus in [#49](https://github.com/outscale/octl/pull/49)

### 🐛 Fixed
* 🐛 fix: fix compilation bug by @jfbus in [#50](https://github.com/outscale/octl/pull/50)

## [v0.0.7] - 2026-02-13

### ✨ Added
* ✨ feat: add delete aliases by @jfbus in [#38](https://github.com/outscale/octl/pull/38)

### 🛠️ Changed / Refactoring
* 💄 ui(confirm): use borders for confirm options by @jfbus in [#41](https://github.com/outscale/octl/pull/41)

### 🐛 Fixed
* 🐛 fix(errors): api errors crashed the output by @jfbus in [#39](https://github.com/outscale/octl/pull/39)

## [v0.0.6] - 2026-02-12

### ✨ Added
* ✨ feat: add OKS provider by @jobs62 in [#34](https://github.com/outscale/octl/pull/34)

### 🛠️ Changed / Refactoring
* 🚸 ux: add some entity aliases/columns definition by @jfbus in [#33](https://github.com/outscale/octl/pull/33)
* 🚸 ux(flags): add required flags by @jfbus in [#36](https://github.com/outscale/octl/pull/36)
* ♻️ change name to octl by @jfbus in [#37](https://github.com/outscale/octl/pull/37)

### 🐛 Fixed
* 🐛 fix(generator): fix iaas api generation by @jfbus in [#35](https://github.com/outscale/octl/pull/35)

## [v0.0.5] - 2026-02-11

### 🛠️ Changed / Refactoring
* 🚸 ux(list): change list shortcut to ls by @jfbus in [#31](https://github.com/outscale/octl/pull/31)
* 🚸 ux(tables): add name tag by @jfbus in [#32](https://github.com/outscale/octl/pull/32)
* 🚸 ux(tables): add colors to states by @jfbus in [#30](https://github.com/outscale/octl/pull/30)

## [v0.0.4] - 2026-02-10

### ✨ Added
* ✨ feat: add high level cli with table output by @jfbus in [#25](https://github.com/outscale/octl/pull/25)

### 🛠️ Changed / Refactoring
* 🚸 ux(slices): dynamically set NumEntriesInSlices by @jfbus in [#22](https://github.com/outscale/octl/pull/22)
* 🚸 ux(profile): add flags for profile file/name by @jfbus in [#23](https://github.com/outscale/octl/pull/23)
* 🚸 ux(templating): add templates by @jfbus in [#24](https://github.com/outscale/octl/pull/24)
* 🚸 ux(times): allow durations or day/month/year offsets in time-based flags by @jfbus in [#26](https://github.com/outscale/octl/pull/26)

### 🐛 Fixed
* 🐛 fix(completion): fix completion of enums in aliases by @jfbus in [#27](https://github.com/outscale/octl/pull/27)
* 🐛 fix(auth): config/profile flags were broken by @jfbus in [#28](https://github.com/outscale/octl/pull/28)

## [v0.0.3] - 2026-02-06

### ✨ Added
* ✨ feat(help): add help texts by @jfbus in [#15](https://github.com/outscale/octl/pull/15)
* ✨ feat(chaining): chain commands by @jfbus in [#18](https://github.com/outscale/octl/pull/18)
* ✨ feat(input): send a raw JSON document by @jfbus in [#19](https://github.com/outscale/octl/pull/19)
* ✨ feat: aliases by @jfbus in #21

### 🛠️ Changed / Refactoring
* ♻️ refacto: move version to --version flag by @jfbus in [#16](https://github.com/outscale/octl/pull/16)
* ♻️ refacto: add builder/runner abstractions by @jfbus in [#20](https://github.com/outscale/octl/pull/20)

### 📝 Documentation
* 📝 doc(readme): improvements by @jfbus in [#11](https://github.com/outscale/octl/pull/11)

## [v0.0.2] - 2026-02-04

### ✨ Added
* ✨ feat(version): add API version by @jfbus in [#12](https://github.com/outscale/octl/pull/12)
* ✨ feat(flags): autocompletion of enums by @jfbus in [#14](https://github.com/outscale/octl/pull/14)

### 🛠️ Changed / Refactoring
* 🚸 ux(booleans): display warning when using --flag false instead of --flag=false by @jfbus in [#13](https://github.com/outscale/octl/pull/13)

## [v0.0.1] - 2026-02-03

### ✨ Added
* 🎉 first version of the cli by @jfbus in [#1](https://github.com/outscale/octl/pull/1)
* ✨ feat(output): add jq-like output filter by @jfbus in [#7](https://github.com/outscale/octl/pull/7)
* ✨ feat(update): add autoupdate by @jfbus in [#8](https://github.com/outscale/octl/pull/8)
### 🐛 Fixed
* 🐛 fix(flags): fix for date slices flags by @jfbus in [#6](https://github.com/outscale/octl/pull/6)
