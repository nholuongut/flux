## 0.14.1 (2019-09-04)

### Improvements

 - Updated Flux to `1.14.2`
   [nholuongut/flux#2419](https://github.com/nholuongut/flux/pull/2419)

## 0.14.0 (2019-08-22)

### Improvements

 - Updated Flux to `1.14.1`
   [nholuongut/flux#2401](https://github.com/nholuongut/flux/pull/2401)
 - Add the ability to disable memcached and set an external memcached service
   [nholuongut/flux#2393](https://github.com/nholuongut/flux/pull/2393)

## 0.13.0 (2019-08-21)

### Improvements

**Note** The Flux chart is now hosted at `https://charts.https://github.com/nholuongut`

 - Updated Flux to `1.14.0`
   [nholuongut/flux#2380](https://github.com/nholuongut/flux/pull/2380)
 - Add `git.readonly` option to chart
   [nholuongut/flux#1807](https://github.com/nholuongut/flux/pull/1807)
 - Helm chart repository has been changed to `charts.https://github.com/nholuongut`
   [nholuongut/flux#2341](https://github.com/nholuongut/flux/pull/2341)

## 0.12.0 (2019-08-08)

### Improvements

 - Updated Flux to `1.13.3` and the Helm operator to `0.10.1`
   [nholuongut/flux#2296](https://github.com/nholuongut/flux/pull/2296)
   [nholuongut/flux#2318](https://github.com/nholuongut/flux/pull/2318)
 - Add manifest generation to helm chart
   [nholuongut/flux#2332](https://github.com/nholuongut/flux/pull/2332)
   [nholuongut/flux#2335](https://github.com/nholuongut/flux/pull/2335)
 - Let a named cluster role be used in chart
   [nholuongut/flux#2266](https://github.com/nholuongut/flux/pull/2266)

## 0.11.0 (2019-07-10)

### Improvements

 - Updated Flux to `1.13.2` and the Helm operator to `0.10.0`
   [nholuongut/flux#2235](https://github.com/nholuongut/flux/pull/2235)
   [nholuongut/flux#2237](https://github.com/nholuongut/flux/pull/2237)
 - Changed from DockerHub organization `nholuongut` -> `nholuongut`
   [nholuongut/flux#2224](https://github.com/nholuongut/flux/pull/2224)
 - Updated `HelmRelease` CRD to support rollbacks
   [nholuongut/flux#2006](https://github.com/nholuongut/flux/pull/2006)
 - Allow namespace scoping for both Flux and the Helm operator
   [nholuongut/flux#2206](https://github.com/nholuongut/flux/pull/2206)
   [nholuongut/flux#2209](https://github.com/nholuongut/flux/pull/2209)
 - Removed long deprecated `FluxHelmRelease` CRD and disabled CRD
   creation as the default to follow our own best practices
   [nholuongut/flux#2190](https://github.com/nholuongut/flux/pull/2190)
 - Enable `PodSecurityPolicy`
   [nholuongut/flux#2223](https://github.com/nholuongut/flux/pull/2223)
   [nholuongut/flux#2225](https://github.com/nholuongut/flux/pull/2225)
 - Support new Flux `--registry-use-labels` flag (`registry.useTimestampLabels`)
   [nholuongut/flux#2176](https://github.com/nholuongut/flux/pull/2176)
 - Support new Helm operator `--workers` flag (`helmOperator.workers`)
   [nholuongut/flux#2236](https://github.com/nholuongut/flux/pull/2236)

## 0.10.2 (2019-06-27)

### Improvements

 - Updated Flux to `1.13.1`
   [nholuongut/flux#2203](https://github.com/nholuongut/flux/pull/2203)

## 0.10.1 (2019-06-16)

### Bug fixes

 - Fix memcached security context
   [nholuongut/flux#2163](https://github.com/nholuongut/flux/pull/2163)

## 0.10.0 (2019-06-14)

### Improvements

 - Updated Flux to `1.13.0` and Helm operator to `0.9.2`
   [nholuongut/flux#2150](https://github.com/nholuongut/flux/pull/2150)
   [nholuongut/flux#2153](https://github.com/nholuongut/flux/pull/2153)
 - Updated memcached to `1.5.15` and configured default security context
   [nholuongut/flux#2107](https://github.com/nholuongut/flux/pull/2107)
 - Toggle garbage collection dry-run
   [nholuongut/flux#2063](https://github.com/nholuongut/flux/pull/2063)
 - Toggle git signature verification
   [nholuongut/flux#2053](https://github.com/nholuongut/flux/pull/2053)
 - Support `dnsPolicy` and `dnsConfig` in Flux daemon deployment
   [nholuongut/flux#2116](https://github.com/nholuongut/flux/pull/2116)
 - Support configurable log format
   [nholuongut/flux#2138](https://github.com/nholuongut/flux/pull/2138)
 - Support additional sidecar containers
   [nholuongut/flux#2130](https://github.com/nholuongut/flux/pull/2130)

### Bug fixes

 - Fix `extraVolumes` indentation
   [nholuongut/flux#2102](https://github.com/nholuongut/flux/pull/2102)

## 0.9.5 (2019-05-22)

 - Updated Flux to `1.12.3`
   [nholuongut/flux#2076](https://github.com/nholuongut/flux/pull/2076)

## 0.9.4 (2019-05-09)

 - Updated Helm operator to `0.9.1`
   [nholuongut/flux#2032](https://github.com/nholuongut/flux/pull/2032)

## 0.9.3 (2019-05-08)

### Improvements

 - Updated Flux to `1.12.2` and Helm operator to `0.9.0`
   [nholuongut/flux#2025](https://github.com/nholuongut/flux/pull/2025)
 - Mount sub path of repositories secret
   [nholuongut/flux#2014](https://github.com/nholuongut/flux/pull/2014)
 - Toggle garbage collection
   [nholuongut/flux#2004](https://github.com/nholuongut/flux/pull/2004)

## 0.9.2 (2019-04-29)

### Improvements

 - Updated Flux to `1.12.1`
   [nholuongut/flux#1993](https://github.com/nholuongut/flux/pull/1993)

## 0.9.1 (2019-04-17)

### Improvements

 - Add the `status` subresource to HelmRelease CRD
   [nholuongut/flux#1906](https://github.com/nholuongut/flux/pull/1906)
 - Switch image registry from Quay to Docker Hub
   [nholuongut/flux#1949](https://github.com/nholuongut/flux/pull/1949)

## 0.9.0 (2019-04-11)

### Improvements

 - Updated Flux to `1.12.0` and Helm operator to `0.8.0`
   [nholuongut/flux#1924](https://github.com/nholuongut/flux/pull/1924)
 - Add ECR require option
   [nholuongut/flux#1863](https://github.com/nholuongut/flux/pull/1863)
 - Support loading values from alternative files in chart 
   [nholuongut/flux#1909](https://github.com/nholuongut/flux/pull/1909)
 - Add Git poll interval option
   [nholuongut/flux#1910](https://github.com/nholuongut/flux/pull/1910)
 - Add init container, extra volumes and volume mounts
   [nholuongut/flux#1918](https://github.com/nholuongut/flux/pull/1918)
 - Add docker config file path option
   [nholuongut/flux#1919](https://github.com/nholuongut/flux/pull/1919)

## 0.8.0 (2019-04-04)

### Improvements

 - Updated Flux to `1.11.1`
   [nholuongut/flux#1892](https://github.com/nholuongut/flux/pull/1892)
 - Define custom Helm repositories in the Helm chart
   [nholuongut/flux#1893](https://github.com/nholuongut/flux/pull/1893)
 - Increase memcached max memory to 512MB
   [nholuongut/flux#1900](https://github.com/nholuongut/flux/pull/1900)

## 0.7.0 (2019-03-27)

### Improvements

 - Updated Flux to `1.11.0` and Helm operator to `0.7.1`
   [nholuongut/flux#1871](https://github.com/nholuongut/flux/pull/1871)
 - Allow mounting of docker credentials file
   [nholuongut/flux#1762](https://github.com/nholuongut/flux/pull/1762)
 - Increase memcached memory defaults
   [nholuongut/flux#1780](https://github.com/nholuongut/flux/pull/1780)
 - GPG Git commit signing
   [nholuongut/flux#1394](https://github.com/nholuongut/flux/pull/1394)

## 0.6.3 (2019-02-14)

### Improvements

 - Updated Flux to `1.10.1`
   [nholuongut/flux#1740](https://github.com/nholuongut/flux/pull/1740)
 - Add option to set pod annotations
   [nholuongut/flux#1737](https://github.com/nholuongut/flux/pull/1737)

## 0.6.2 (2019-02-11)

### Improvements

 - Allow chart images to be pulled from a private container registry
   [nholuongut/flux#1718](https://github.com/nholuongut/flux/pull/1718)

### Bug fixes

 - Fix helm-op allow namespace flag mapping
   [nholuongut/flux#1724](https://github.com/nholuongut/flux/pull/1724)

## 0.6.1 (2019-02-07)

### Improvements

 - Updated Flux to `1.10.0` and Helm operator to `0.6.0`
   [nholuongut/flux#1713](https://github.com/nholuongut/flux/pull/1713)
 - Add option to exclude container images
   [nholuongut/flux#1659](https://github.com/nholuongut/flux/pull/1659)
 - Add option to mount custom `repositories.yaml`
   [nholuongut/flux#1671](https://github.com/nholuongut/flux/pull/1671)
 - Add option to limit the Helm operator to a single namespace
   [nholuongut/flux#1664](https://github.com/nholuongut/flux/pull/1664)

### Bug fixes

 - Fix custom SSH secret mapping
   [nholuongut/flux#1710](https://github.com/nholuongut/flux/pull/1710)

## 0.6.0 (2019-01-14)

**Note** To fix the connectivity problems between Flux and memcached we've changed the
memcached service from headless to ClusterIP. This change will make the Helm upgrade fail
with `ClusterIP field is immutable`.

Before upgrading to 0.6.0 you have to delete the memcached headless service:

```bash
kubectl -n flux delete svc flux-memcached
```

### Improvements

 - Updated Flux to `1.9.0` and Helm operator to `0.5.3`
   [nholuongut/flux#1662](https://github.com/nholuongut/flux/pull/1662)
 - Add resetValues field to HelmRelease CRD
   [nholuongut/flux#1628](https://github.com/nholuongut/flux/pull/1628)
 - Use ClusterIP service name for connecting to memcached
   [nholuongut/flux#1618](https://github.com/nholuongut/flux/pull/1618)
 - Increase comprehensiveness of values table in `chart/flux/README.md`
   [nholuongut/flux#1626](https://github.com/nholuongut/flux/pull/1626)
    - Rectify error where `resources` are not `None` by default in `chart/flux/values.yaml`
    - Add more fields that are actually in `chart/flux/values.yaml`
    - Separate `replicaCount` into a Flux one and `helmOperator.replicaCount` one
  - Only create the `flux-helm-tls-ca-config` file if `.Values.helmOperator.tls.caContent` exists.
    Useful when doing Flux upgrades but do not happen to know or want to specify
    the `caContent` in `values.yaml`. Otherwise, the existing caContent will be overriden with an
    empty value.
    [nholuongut/flux#1649](https://github.com/nholuongut/flux/pull/1649)
  - Add Flux AWS ECR flags
    [nholuongut/flux#1655](https://github.com/nholuongut/flux/pull/1655)


## 0.5.2 (2018-12-20)

### Improvements

 - Updated Flux to `v1.8.2` and Helm operator to `v0.5.2`
   [nholuongut/flux#1612](https://github.com/nholuongut/flux/pull/1612)
 - Parameterized the memcached image repo
   [nholuongut/flux#1592](https://github.com/nholuongut/flux/pull/1592)
 - Allow existing service account to be provided on helm install
   [nholuongut/flux#1589](https://github.com/nholuongut/flux/pull/1589)
 - Make SSH known hosts volume optional
   [nholuongut/flux#1544](https://github.com/nholuongut/flux/pull/1544)

### Thanks

Thanks to @davidkarlsen, @stephenmoloney, @batpok, @squaremo,
@hiddeco and @stefanprodan for their contributions.

## 0.5.1 (2018-11-21)

### Bug fixes

 - Removed CRD hook from chart
   [nholuongut/flux#1536](https://github.com/nholuongut/flux/pull/1536)

### Improvements

 - Updated Helm operator to `v0.5.1`
   [nholuongut/flux#1536](https://github.com/nholuongut/flux/pull/1536)
 - Updated chart README (removed Helm operator Git flags, fixed typos,
   updated example repo and use the same Git URL format everywhere)
   [nholuongut/flux#1527](https://github.com/nholuongut/flux/pull/1527)

## 0.5.0 (2018-11-16)

### Improvements

 - Updated Flux to `v1.8.1` and the Helm operator to `v0.5.0`
   [nholuongut/flux#1522](https://github.com/nholuongut/flux/pull/1522)
 - Adapted chart to new Helm operator CRD and args
   [nholuongut/flux#1382](https://github.com/nholuongut/flux/pull/1382)

## 0.4.1 (2018-11-04)

### Bug fixes

 - Fixed indentation of `.Values.helmOperator.tls.caContent`
   [nholuongut/flux#1484](https://github.com/nholuongut/flux/pull/1484)

### Improvements

 - Updated Helm operator to `v0.4.0`
   [nholuongut/flux#1487](https://github.com/nholuongut/flux/pull/1487)
 - Added `--tiller-tls-hostname` Helm operator config flag to the chart
   [nholuongut/flux#1484](https://github.com/nholuongut/flux/pull/1484)
 - Include `valueFileSecrets` property in `helm-operator-crd.yaml`
   [nholuongut/flux#1468](https://github.com/nholuongut/flux/pull/1468)
 - Uniform language highlight on Helm chart README
   [nholuongut/flux#1464](https://github.com/nholuongut/flux/pull/1463)

## 0.4.0 (2018-10-25)

### Bug fixes

 - Made maximum memcache item size configurable, fixes
   `SERVER_ERROR object too large for cache`  errors on large deployments
   [nholuongut/flux#1453](https://github.com/nholuongut/flux/pull/1453)
 - Fixed indentation of `aditionalArgs`
   [nholuongut/flux#1417](https://github.com/nholuongut/flux/pull/1417)

### Improvements

 - Updated Flux to `v1.8.0` and the Helm operator to `0.3.0`
   [nholuongut/flux#1470](https://github.com/nholuongut/flux/pull/1470)
 - Deprecated Flux `--registry-cache-expiry` config flag
   [nholuongut/flux#1470](https://github.com/nholuongut/flux/pull/1470)
 - Added and documented multiple values (s.a. `nodeSelector`,
   `extraEnvs`, `git.timeout`)
   [nholuongut/flux#1469](https://github.com/nholuongut/flux/pull/1469)
   [nholuongut/flux#1446](https://github.com/nholuongut/flux/pull/1446)
   [nholuongut/flux#1416](https://github.com/nholuongut/flux/pull/1416)
 - Made it possible to enable Promotheus annotations
   [nholuongut/flux#1462](https://github.com/nholuongut/flux/pull/1462)

## 0.3.4 (2018-09-28)

### Improvements

 - Updated Flux to `v1.7.1`
   [nholuongut/flux#1405](https://github.com/nholuongut/flux/pull/1405)
 - Custom SSH keys for Flux and Helm operator are now allowed
   [nholuongut/flux#1391](https://github.com/nholuongut/flux/pull/1391)

## 0.3.3 (2018-09-18)

### Improvements

 - Updated Flux to `v1.7.0` and the Helm operator to `v0.2.1`
   [nholuongut/flux#1368](https://github.com/nholuongut/flux/pull/1368)
 - Added memcached verbose option
   [nholuongut/flux#1350](https://github.com/nholuongut/flux/pull/1350)
 - Allow overrides of `.kube/config`
   [nholuongut/flux#1342](https://github.com/nholuongut/flux/pull/1342)
 - Documentation improvements
   [nholuongut/flux#1357](https://github.com/nholuongut/flux/pull/1357)

## 0.3.2 (2018-08-31)

### Improvements

 - Updated Flux to `v1.6.0`
   [nholuongut/flux#1330](https://github.com/nholuongut/flux/pull/1330)
 - Made the Helm operator CRD creation optional
   [nholuongut/flux#1311](https://github.com/nholuongut/flux/pull/1311)

## 0.3.0 (2018-08-24)

### Improvements

 - Updated Helm operator to `v0.2.0`
   [nholuongut/flux#1308](https://github.com/nholuongut/flux/pull/1308)
 - Added Flux git label and registry options
   [nholuongut/flux#1305](https://github.com/nholuongut/flux/pull/1305)
 - Removed `.Values.git.gitPath` value
   [nholuongut/flux#1305](https://github.com/nholuongut/flux/pull/1305)
 - Documented how to use a private Git host
   [nholuongut/flux#1299](https://github.com/nholuongut/flux/pull/1299)
 - Added option to opt-in to logging of release diffs
   [nholuongut/flux#1271](https://github.com/nholuongut/flux/pull/1272)

## 0.2.2 (2018-08-09)

### Bug fixes

 - Fixed indentation of `.Values.ssh.known_hosts`
   [nholuongut/flux#1246](https://github.com/nholuongut/flux/pull/1246)

### Improvements

 - Updated Flux to `v1.5.0`
   [nholuongut/flux#1279](https://github.com/nholuongut/flux/pull/1279)
 - Added openAPIV3Schema validation to Helm CRD
   [nholuongut/flux#1253](https://github.com/nholuongut/flux/pull/1253)
 - Fix markdown typo in README
   [nholuongut/flux#1248](https://github.com/nholuongut/flux/pull/1248)
