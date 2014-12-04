{{ define "Content" }}
<div ng-app="pubApp" ng-controller="EditController">
<div class="row">
	<div class="col-lg-12">
		<h3>Edit Post</h3>
	</div>
</div>
<div class="row">
	<div class="col-lg-9">
		<input type="text" ng-model="post.title" class="form-control" />
		<br />
		<div ui-ace ng-model="post.content"></div>
	</div>
	<div class="col-lg-3">
		<button ng-click="savePost();" type="submit" style="width: 100%;" class="btn btn-primary">Save</button>
		<br />
		<br />
		<div class="panel panel-default">
            <div class="panel-heading">
                <h3 class="panel-title"><i class="fa fa-fw fa-file-text-o"></i> Publish</h3>
            </div>
            <div class="panel-body">
				<input type="checkbox" ng-click="publishClick()" ng-model="post.published" /> Publish
            </div>
        </div>
        <div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Category</h3>
			</div>
			<div class="panel-body">
				<select ng-model="post.category" ng-options="c.title for c in cats track by c.id">
					<option value="">-- choose category --</option>
				</select>
			</div>
        </div>
	</div>
</div>
</div>
{{ end }}

{{ define "Scripts" }}
<script language="JavaScript">
	var backend = {};
	backend.post = {{ marshal .post }};
	{{ if .cats }}
	backend.cats = {{ marshal .cats }};
	{{ end }}
</script>
<script src="/admin/js/ace-builds/src-min-noconflict/ace.js"></script>
<script src="/admin/js/angular.js"></script>
<script src="/admin/js/ui-ace.js"></script>
<script src="/admin/js/angular-resource.js"></script>
<script src="/admin/js/apps/publishApp.js"></script>
<script src="/admin/js/apps/editApp.js"></script>
{{ end }}