# Changelog

## 0.1.0 (2026-05-14)

Full Changelog: [v0.0.3...v0.1.0](https://github.com/et0and/schools-sdk-go/compare/v0.0.3...v0.1.0)

### Features

* **go:** add default http client with timeout ([c7040a9](https://github.com/et0and/schools-sdk-go/commit/c7040a95e2c5a7a375e5251797c0b43a49cb4a90))
* **internal:** support comma format in multipart form encoding ([4dd56d3](https://github.com/et0and/schools-sdk-go/commit/4dd56d322dbf3fa7497c0c42470407ef23b945bd))
* support setting headers via env ([c7506bd](https://github.com/et0and/schools-sdk-go/commit/c7506bd714ea79a7668271edbb80a3241ebc810c))


### Bug Fixes

* prevent duplicate ? in query params ([458cdb9](https://github.com/et0and/schools-sdk-go/commit/458cdb9e3eda53926b5571c3784d124055b399fc))


### Chores

* **ci:** skip lint on metadata-only changes ([15e379e](https://github.com/et0and/schools-sdk-go/commit/15e379e4aca26266007227d761ab13860140a156))
* **ci:** skip uploading artifacts on stainless-internal branches ([dd0dab4](https://github.com/et0and/schools-sdk-go/commit/dd0dab4dc3822344b18e2f6a27f0ad781e7aca12))
* **ci:** support opting out of skipping builds on metadata-only commits ([fdd5897](https://github.com/et0and/schools-sdk-go/commit/fdd589734623922da30c723df0e7240f2ecce969))
* **client:** fix multipart serialisation of Default() fields ([688c492](https://github.com/et0and/schools-sdk-go/commit/688c492ffa8d0df8ad27cb707d5f83e33f414399))
* **internal:** codegen related update ([b8ead58](https://github.com/et0and/schools-sdk-go/commit/b8ead586a6ab6138fc6e0ce698240e92f91b3326))
* **internal:** codegen related update ([5011a1f](https://github.com/et0and/schools-sdk-go/commit/5011a1f1c56755d7e4886164b8e01678b05379e3))
* **internal:** codegen related update ([b9eb07b](https://github.com/et0and/schools-sdk-go/commit/b9eb07b6db5be3b5d171d05d9f2db81464dee90e))
* **internal:** codegen related update ([930d631](https://github.com/et0and/schools-sdk-go/commit/930d6311b534d5dde11975202c56e5a92effced8))
* **internal:** codegen related update ([b94d863](https://github.com/et0and/schools-sdk-go/commit/b94d863327ac7cbc15a274b81cf7f32fc8b8f73e))
* **internal:** codegen related update ([be3d26c](https://github.com/et0and/schools-sdk-go/commit/be3d26cacec89f0eedf13ac24e874cb34b6d7dec))
* **internal:** minor cleanup ([5e6a0c9](https://github.com/et0and/schools-sdk-go/commit/5e6a0c9d4b9bc890e6158f9ab47f4f2e2fee411a))
* **internal:** more robust bootstrap script ([ea06030](https://github.com/et0and/schools-sdk-go/commit/ea06030ec572a4fe34ffea04474b0bb63074a551))
* **internal:** support default value struct tag ([f232335](https://github.com/et0and/schools-sdk-go/commit/f232335e60cad792c6d196b4bfb25630414922d7))
* **internal:** tweak CI branches ([3232046](https://github.com/et0and/schools-sdk-go/commit/3232046f134581346788fb3a3bf781d9377462cb))
* **internal:** update gitignore ([53cb932](https://github.com/et0and/schools-sdk-go/commit/53cb93296a28222746148ae8c20c81b02718bdc7))
* **internal:** use explicit returns ([72510a5](https://github.com/et0and/schools-sdk-go/commit/72510a5c3b3a5d1c9e70b9cadc6b5c773e530790))
* **internal:** use explicit returns in more places ([3e9528b](https://github.com/et0and/schools-sdk-go/commit/3e9528ba7865ecee2a1dff8f90464f85674596b9))
* remove unnecessary error check for url parsing ([9534637](https://github.com/et0and/schools-sdk-go/commit/953463779c6e90dd15df21722fa6ceec8acabb5a))
* update docs for api:"required" ([920f779](https://github.com/et0and/schools-sdk-go/commit/920f7794160fa9bd6bdfd6f0657740bf628f48fb))

## 0.0.3 (2026-02-25)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/et0and/schools-sdk-go/compare/v0.0.2...v0.0.3)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([d620f77](https://github.com/et0and/schools-sdk-go/commit/d620f77b7ca03538c68d79a8a05cbead1fbc0113))
* **encoder:** correctly serialize NullStruct ([3971a18](https://github.com/et0and/schools-sdk-go/commit/3971a187e4ca2f265bc2fb6e4fda2822fcb944d2))


### Chores

* **internal:** move custom custom `json` tags to `api` ([08c5ca3](https://github.com/et0and/schools-sdk-go/commit/08c5ca3b8bccd6f9f3f4c678e9b8819cdcc13bc6))
* **internal:** remove mock server code ([2961b8b](https://github.com/et0and/schools-sdk-go/commit/2961b8bc35092d52777937279ce47ddf373041a4))
* update mock server docs ([bdd85b7](https://github.com/et0and/schools-sdk-go/commit/bdd85b79c9f10d767f4d6395fa26c80b551c750a))

## 0.0.2 (2026-01-30)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/et0and/schools-sdk-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([9fc68bc](https://github.com/et0and/schools-sdk-go/commit/9fc68bcc34b20189d5c8bd3f1cf0fdb1e57edd5b))
* update SDK settings ([e270df9](https://github.com/et0and/schools-sdk-go/commit/e270df92686fdfa148f7bda33e02bfdec6fc95fc))
