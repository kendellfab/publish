{{ define "Content" }}
<div ng-app="pubApp" ng-controller="EditController">
<div class="row">
	<div class="col-lg-9">
		<h3>Edit Post</h3>
		<input type="text" ng-model="post.title" class="form-control" />
		<textarea ng-model="post.content" cols="80" rows="15"></textarea>
	</div>
	<div class="col-lg-3">
		<button ng-click="savePost();" type="submit" class="btn btn-primary">Save</button>
		<div class="panel panel-default">
            <div class="panel-heading">
                <h3 class="panel-title"><i class="fa fa-fw fa-file-text-o"></i> Publish</h3>
            </div>
            <div class="panel-body">

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
</script>
<script src="/admin/js/angular.js"></script>
<script src="/admin/js/angular-resource.js"></script>
<script src="/admin/js/apps/publishApp.js"></script>
<script src="/admin/js/apps/editApp.js"></script>
{{ end }}