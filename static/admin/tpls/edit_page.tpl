{{ define "Content" }}
<div ng-app="pubApp" ng-controller="PageController">
<div class="row">
	<div class="col-lg-12">
		<h3>Edit Page</h3>
	</div>
</div>
<div class="row">
	<div class="col-lg-9">
		<input type="text" ng-model="page.title" class="form-control" />
		<br />
		<div ui-ace="{useWrapMode: true, mode: 'markdown'}" ng-model="page.content"></div>
	</div>
	<div class="col-lg-3">
		<button ng-click="savePage();" type="submit" style="width: 100%;" class="btn btn-primary">Save</button>
		<br />
		<br />
		<div ng-show="saving" class="save-indicator">
			<img src="/admin/img/ajax-loader.gif" /> <span>Saving...</span>
		</div>
		<div class="panel panel-default">
            <div class="panel-heading">
                <h3 class="panel-title"><i class="fa fa-fw fa-file-text-o"></i> Publish</h3>
            </div>
            <div class="panel-body">
				<input type="checkbox" ng-click="publishClick()" ng-model="page.published" /> Publish
            </div>
        </div>
        <div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Delete</h3>
			</div>
			<div class="panel-body">
				<a href="/admin/pages/{{ .page.Id }}/delete" class="btn btn-danger full">Delete</a>
			</div>
        </div>
	</div>
</div>
</div>
{{ end }}

{{ define "Scripts" }}
<script language="JavaScript">
	var backend = {};
	backend.page = {{ marshal .page }};
</script>
<script src="/admin/js/ace-builds/src-min-noconflict/ace.js"></script>
<script src="/admin/js/angular.js"></script>
<script src="/admin/js/ui-ace.js"></script>
<script src="/admin/js/angular-resource.js"></script>
<script src="/admin/js/apps/publishApp.js"></script>
<script src="/admin/js/apps/pageApp.js"></script>
{{ end }}