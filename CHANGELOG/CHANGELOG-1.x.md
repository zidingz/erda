# Erda Changelog 1.x

<table>
<tr>
  <th title="Current">1.2<sup>Current</sup></th>
  <th title="Current">1.1</th>
  <th title="Current">1.0</th>
</tr>
<tr>
  <td valign="top">
    <b><a href="#v121">1.2.1</a></b><br/>
    <a href="#v120">1.2.0</a><br/>
  </td>
  <td valign="top">
    <a href="#v111">1.1.1</a><br/>
    <a href="#v110">1.1.0</a><br/>
  </td>
  <td valign="top">
    <a href="#v101">1.0.1</a><br/>
    <a href="#v100">1.0.0</a><br/>
  </td>
</tr>
</table>

# v1.2.1

### Bug Fixes

`2021-08-23`

* Fix the issue of table style for manual test case. [#953](https://github.com/erda-project/erda-ui/pull/953)
* Fix the bug that occured when switching the source type for the first time after resetting the form while adding tags. [#957](https://github.com/erda-project/erda-ui/pull/957)
* Fix the bug of regular expression of repository address field. [#958](https://github.com/erda-project/erda-ui/pull/958)
* Fix the issue of purple label without background color. [#962](https://github.com/erda-project/erda-ui/pull/962)
* Fix the bug that the empty page icon is not displayed when there is no branch on the API design page. [#970](https://github.com/erda-project/erda-ui/pull/970)
* Fix the bug of tracking details type.[#975](https://github.com/erda-project/erda-ui/pull/975)
* Fix the bug that cluster_name and application_id do not take effect when they exist in custom filter rules. [#1459](https://github.com/erda-project/erda/pull/1459)
* Fix the issue that the ES index of log analysis is not scrolling.[#1464](https://github.com/erda-project/erda/pull/1464)[#1465](https://github.com/erda-project/erda/pull/1465)
* Fix the issue of memory leak when getting the instance list.[#1493](https://github.com/erda-project/erda/pull/1493)
* Support getting the specified pod when obtaining the pod status list.[#1495](https://github.com/erda-project/erda/pull/1495)

# v1.2.0

`2021-08-16`

### New Features

* Optimize overview and project list in MSP. [#809](https://github.com/erda-project/erda-ui/pull/809)
* Support sending test messages when configuring DingTalk notifications. [#777](https://github.com/erda-project/erda-ui/pull/777)
* Support importing and exporting automation test sets. [#749](https://github.com/erda-project/erda-ui/pull/749)
* Enable more features of multi-cloud management platform for free users. [#759](https://github.com/erda-project/erda-ui/pull/759)
* Optimize the way to add EDAS clusters. [#750](https://github.com/erda-project/erda-ui/pull/750)
* Optimize markdown editor interaction and style. [#853](https://github.com/erda-project/erda-ui/pull/853)
* Optimize pipeline log style. [#802](https://github.com/erda-project/erda-ui/pull/802)
* Optimize pipeline notification content. [#1189](https://github.com/erda-project/erda/pull/1189)
* Optimize the execution logic of automation test cases. [#1103](https://github.com/erda-project/erda/pull/1103)
* Support filtering test case executor by unassigned person in the test plan. [#732](https://github.com/erda-project/erda/pull/732)
* Add precheck for existence of package-lock.json before packaging front-end applicaiton. [#1116](https://github.com/erda-project/erda/pull/1116)

### Bug Fixes

* Fix a bug of cluster parameter in the project pipeline. [#737](https://github.com/erda-project/erda-ui/pull/737)
* Fix the bug of data duplication in repo file comparison. [#744](https://github.com/erda-project/erda-ui/pull/744)
* Modify markdown editor style. [#763](https://github.com/erda-project/erda-ui/pull/763)
* Fix a bug of env parameter in the project pipeline. [#765](https://github.com/erda-project/erda-ui/pull/765)
* Fix the style issue when dragging and droppping Nusi component tree. [#769](https://github.com/erda-project/erda-ui/pull/769)
* Fix the error of operation key value in action form. [#771](https://github.com/erda-project/erda-ui/pull/771)
* Fix the mandatory verification error of custom labels in the form. [#778](https://github.com/erda-project/erda-ui/pull/778)
* Fix the button style issue of markdown editor. [#782](https://github.com/erda-project/erda-ui/pull/782)
* Fix the issue of yml editor add node disappearance and actionForm componentization data error. [#781](https://github.com/erda-project/erda-ui/pull/781)
* Fix the bug of the drop-down width when selecting artifact ID in deployment center.[#827](https://github.com/erda-project/erda-ui/pull/827)
* Add registration command and retry initialization operations for EDAS cluster. [#840](https://github.com/erda-project/erda-ui/pull/840)
* Fix the bug of extension service form in project. [#863](https://github.com/erda-project/erda-ui/pull/863)
* Fix the possible crash bug when adding members using nicknames with special characters. [#862](https://github.com/erda-project/erda-ui/pull/862)
* Fix the bug that projectId is missing in the request application list. [#873](https://github.com/erda-project/erda-ui/pull/873)
* Fix the bug that Git repositories can be cloned without account password.[#877](https://github.com/erda-project/erda-ui/pull/877)
* Fix the bug that the text prompt is incomplete caused by invalid form in trace debugging. [#857](https://github.com/erda-project/erda-ui/pull/857)
* Fix the data error of related issues after changing issue to backlog. [#902](https://github.com/erda-project/erda-ui/pull/902)
* Fix the bug that two scroll bars appear when scrolling the item in backlog table. [#839](https://github.com/erda-project/erda-ui/pull/839)
* Fix the API error reported after deleting files in repo. [#910](https://github.com/erda-project/erda-ui/pull/910)
* Fix the error occured when initializing pipeline action form. [#912](https://github.com/erda-project/erda-ui/pull/912)
* Fix some table issues: column width too long or insufficient, table exceeds the page and uniform overflow omission. [#736](https://github.com/erda-project/erda-ui/pull/736)
* Fix the style issue that the item name of extended query column is too long. [#739](https://github.com/erda-project/erda-ui/pull/739)
* Fix the bug that required fields in project collaboration are not marked as required.[#746](https://github.com/erda-project/erda-ui/pull/746)
* Add width to the table in OrgCenter > Projects. [#755](https://github.com/erda-project/erda-ui/pull/755)
* Fix the bug of style validation occured when adding issue in Issues > Backlog. [#757](https://github.com/erda-project/erda-ui/pull/757)
* Fix the bug that in Multi-Cloud Management Platform > Alarm Record, click a record for details, then all list items are expanded when clicking the expand button before the list items.[#754](https://github.com/erda-project/erda-ui/pull/754)
* Fix the bug that when editing issues, the month in datepicker cannot be changed.[#761](https://github.com/erda-project/erda-ui/pull/761)
* Increase the width of the member table. [#767](https://github.com/erda-project/erda-ui/pull/767)
* Fix the bug that the row representing the folder in the test case table shows an extra column for the interface pass rate. [#768](https://github.com/erda-project/erda-ui/pull/768)
* Increase the width of the test case table. [#774](https://github.com/erda-project/erda-ui/pull/774)
* Add a mouse-over style to tables with row click events. [#766](https://github.com/erda-project/erda-ui/pull/766)
* Adjust the width of the related issue table. [#775](https://github.com/erda-project/erda-ui/pull/775)
* Fix the loop refresh issue when the path is /-. [#780](https://github.com/erda-project/erda-ui/pull/780)
* Fix the bug that the left arrow faces wrongly when the tree on the left side of the test case is expanded, and the parent node occasionally collapses when the child node is clicked.[#773](https://github.com/erda-project/erda-ui/pull/773)
* Fix the bug that some rows in the test case table go beyond the right side. [#790](https://github.com/erda-project/erda-ui/pull/790)
* Fix the incorrect address of application repository in application settings. [#797](https://github.com/erda-project/erda-ui/pull/797)
* Fix the issue that the color of alert list icon is black. [#808](https://github.com/erda-project/erda-ui/pull/808)
* Fix the issue that error occured when editing custom addon of extended service. [#812](https://github.com/erda-project/erda-ui/pull/812)
* Fix the issue that the text field of trace debugging body is too long to see the tabs above. [#820](https://github.com/erda-project/erda-ui/pull/820)
* Fix the bug that logs cannot be downloaded when using the default duration. [#842](https://github.com/erda-project/erda-ui/pull/842)
* Fix the bug that the search box does not display when the data is empty. [#906](https://github.com/erda-project/erda-ui/pull/906)
* Migrate Affix and InputNumber from Nusi to Antd.[#779](https://github.com/erda-project/erda-ui/pull/779)
* Change the grouping rules to mandatory when creating custom alarm rules. [#752](https://github.com/erda-project/erda-ui/pull/752)
* Fix the issue that the setting of runtime rollback number does not take effect.

### Refactor

* The interfaces of msp and monitor modules are all defined by proto.
* Optimize Quick-Start for one-click deployment of Erda Standalone mode on the local machine. [#1242](https://github.com/erda-project/erda/pull/1242)

# v1.1.1

`2021-8-5`

### Bug Fixes

* EDAS and K8S use the same agent now. ([#1277](https://github.com/erda-project/erda/pull/1277))
* Enable cloud management platform for free users. ([#810](https://github.com/erda-project/erda-ui/pull/810))
* Fixed the issue that error occured when editing custom addon of extended service. ([#813](https://github.com/erda-project/erda-ui/pull/813))

# v1.1.0

`2021-07-28`

### New Features

- Support existing clusters importing by users ([#806](https://github.com/erda-project/erda/pull/806))
- Support subscription to issue changes, to receive notifications timely when followed issue is modified ([#451](https://github.com/erda-project/erda-ui/pull/451))
- Support asynchronous import and export of manual test cases ([#380](https://github.com/erda-project/erda-ui/pull/380))
- Support auto page refresh for automated test plan ([#446](https://github.com/erda-project/erda-ui/pull/446))
- Support bug closed time viewing and filtering ([#445](https://github.com/erda-project/erda/pull/445))
- Add project-level applications to realize rapid project migration ([#350](https://github.com/erda-project/erda-ui/pull/350))
- Support page turning in Issues > Backlog ([#395](https://github.com/erda-project/erda-ui/pull/395))
- Optimize clone address of code repository ([#155](https://github.com/erda-project/erda-ui/pull/155))
- Optimize size of sliding window and description area for issue editing ([#314](https://github.com/erda-project/erda-ui/pull/314))
- Support size adjusting for table pagination ([#1031](https://github.com/erda-project/erda/pull/1031))
- Turn enter search to delayed auto search for personal dashboard ([#324](https://github.com/erda-project/erda-ui/pull/324))
- Optimize the downloaded file name and suffix format of container log: service name_timestamp.log ([#684](https://github.com/erda-project/erda/pull/684/files))
- Upgrade logo ([#688](https://github.com/erda-project/erda-ui/pull/688))

### Bug Fixes

- Safari page crashes when access Code Repository > Commit History ([#384](https://github.com/erda-project/erda-ui/pull/384))
- The list order remains unchanged after viewing MR ([#661](https://github.com/erda-project/erda/pull/661))
- The certificate file is uploaded but its name is not displayed ([#639](https://github.com/erda-project/erda-ui/pull/639))
- Canceling the edit of merge request will clear the comparison result ([#638](https://github.com/erda-project/erda-ui/pull/638))
- Failed to download files in code repository ([#588](https://github.com/erda-project/erda-ui/pull/588))
- The pipeline node shows the previously failed error ([#422](https://github.com/erda-project/erda-ui/pull/422))
- Members will automatically log out if exit the organization ([#347](https://github.com/erda-project/erda-ui/pull/347))

### Refactor

- Split out new platform services of dop, msp, cmp, ecp and admin
  - dop ([#392](https://github.com/erda-project/erda-ui/pull/392))
  - msp ([#407](https://github.com/erda-project/erda-ui/pull/407))
  - cmp ([#416](https://github.com/erda-project/erda-ui/pull/416))
  - ecp ([#419](https://github.com/erda-project/erda-ui/pull/419))
- Remove components of qa, apim, cmdb, ops and tmc
- Add core components of cluster manager
- Add cluster-dialer instead of soldier to handle inter-cluster communication
- Add a new way to define an interface using protobuf protocol, and the msp component has been migrated

# v1.0.1

`2021-07-08`

### New features

- Support individuals to create organization if he/she doesn't belong to any organizations ([#592](https://github.com/erda-project/erda/pull/592))
- Support scaling application without restarting the existing instances ([#644](https://github.com/erda-project/erda/pull/644)) ([#645](https://github.com/erda-project/erda/pull/645))

### Bug Fixes

- Fix the issue that custom stages were created with null value ([#588](https://github.com/erda-project/erda/pull/588)) ([#606](https://github.com/erda-project/erda/pull/606))
- Keep the creator, assignee, create time and man hour unchanged when issue type is switched ([#610](https://github.com/erda-project/erda/pull/610)) ([#612](https://github.com/erda-project/erda/pull/612))
- Fix the logic error of job deletion under the specified namespace ([#632](https://github.com/erda-project/erda/pull/632)) ([#636](https://github.com/erda-project/erda/pull/636))
- Add guest permissions for dashboard and ticket ([#701](https://github.com/erda-project/erda/pull/701)) ([#705](https://github.com/erda-project/erda/pull/705))
- Reset flags before loop for wait step when timed out ([#715](https://github.com/erda-project/erda/pull/715))

# v1.0.0

`2021-06-09`

Erda v1.0.0 is released!

Start your Erda journey in two ways:

- [Quick start in your local machine](https://github.com/erda-project/erda/blob/master/docs/guides/quickstart/quickstart-full.md)
- [Install with K8s](https://github.com/erda-project/erda/blob/master/docs/guides/deploy/How-to-install-Erda.md)
