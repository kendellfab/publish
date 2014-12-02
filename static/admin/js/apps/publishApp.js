var pubApp = angular.module("pubApp", ["ngResource", "ui.ace"]).config(function($interpolateProvider) {
	$interpolateProvider.startSymbol('||');
	$interpolateProvider.endSymbol('||');
});

pubApp.factory('res', ['$resource', function($resource){
	var postResource = $resource("/admin/post/:id/edit", {id: "@id"}, {
		update: {
			method: "Post"
		}
	});

	return {
		post: postResource
	}
}])