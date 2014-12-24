{{ define "Content" }}
<div ng-app="pubApp" ng-controller="SeriesController">
<div class="row">
	<div class="col-lg-12">
		<h3>Edit Series</h3>
	</div>
</div>
<div class="row">
	<div class="col-lg-9">
		<input type="text" ng-model="series.title" class="form-control" />
		<br />
		<div ui-ace="{useWrapMode: true, mode: 'markdown'}" ng-model="series.description"></div>
	</div>
	<div class="col-lg-3">
		<button ng-click="saveSeries();" type="submit" style="width: 100%;" class="btn btn-primary">Save</button>
		<br />
		<br />
		<div ng-show="saving" class="save-indicator">
			<img src="/admin/img/ajax-loader.gif" /> <span>Saving...</span>
		</div>

		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Posts</h3>
			</div>
			<div class="panel-body">
				<table class="table table-bordered table-hover table-striped">
					<tbody>
						<tr ng-repeat="p in series.posts">
							<td><a href="/admin/post/||p.id||/edit">||p.title||</a></td>
							<td><a href="#" ng-click="removePost($index)">&times;</a></td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</div>
</div>
{{ end }}

{{ define "Scripts" }}
<script language="JavaScript">
	var backend = {};
	backend.series = {{ marshal .series }};
</script>
<script src="/admin/js/ace-builds/src-min-noconflict/ace.js"></script>
<script src="/admin/js/angular.js"></script>
<script src="/admin/js/ui-ace.js"></script>
<script src="/admin/js/angular-resource.js"></script>
<script src="/admin/js/apps/publishApp.js"></script>
<script src="/admin/js/apps/seriesApp.js"></script>
{{ end }}