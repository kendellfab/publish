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
		<div ui-ace="{useWrapMode: true, mode: 'markdown'}" ng-model="post.content"></div>
	</div>
	<div class="col-lg-3">
		<button ng-click="savePost();" type="submit" style="width: 100%;" class="btn btn-primary">Save</button>
		<br />
		<br />
		<div ng-show="saving" class="save-indicator">
			<img src="/admin/img/ajax-loader.gif" /> <span>Saving...</span>
		</div>
		<button type="button" class="btn btn-success" style="width: 100%;" data-toggle="modal" data-target="#myModal">Uploads</button>
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
					
				</select>
			</div>
        </div>
        <div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Series</h3>
			</div>
			<div class="panel-body">
				<select ng-model="post.seriesId" ng-options="s.id as s.title for s in series">
					<option value="">-- choose series --</option>
				</select>
			</div>
        </div>
	</div>
</div>

<div class="modal fade" id="myModal">
  <div class="modal-dialog modal-admin">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
        <h4 class="modal-title">Upload List</h4>
      </div>
      <div class="modal-body modal-scroll">
        <table class="table table-bordered table-hover table-striped">
			<thead>
				<tr>
					<th>Thumb</th>
					<th>Name</th>
					<th>Link</th>
				</tr>
			</thead>
			<tbody>
				<tr ng-repeat="file in uploads">
					<td><img src="/uploads/||file.name||" height="50" /></td>
					<td>||file.name||</td>
					<td>/uploads/||file.name||</td>
				</tr>
			</tbody>
        </table>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

</div>
{{ end }}

{{ define "Scripts" }}
<script language="JavaScript">
	var backend = {};
	backend.post = {{ marshal .post }};
	backend.series = {{ marshal .series }};
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