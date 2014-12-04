var pubApp = angular.module("pubApp", ["ngResource", "ui.ace"]).config(function($interpolateProvider) {
	$interpolateProvider.startSymbol('||');
	$interpolateProvider.endSymbol('||');
});

pubApp.factory('res', ['$resource', function($resource){
	var postResource = $resource("/admin/post/:id/:action", {id: "@id"}, {
		update: {
			method: "Post",
			params: {action: "edit"}
		},
		publish: {
			method: "Post",
			params: {action: "publish"}
		}
	});

	return {
		post: postResource
	}
}])