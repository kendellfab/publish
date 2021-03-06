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

	var pageResource = $resource("/admin/pages/:id/:action", {id: "@id"}, {
		update: {
			method: "Post",
			params: {action: "edit"}
		},
		publish: {
			method: "Post",
			params: {action: "publish"}
		}
	})

	var uploadResource = $resource("/admin/uploads/:id/:action", {}, {
		list: {
			method: "Get",
			params: {id: "list"},
			isArray: true
		}
	});

	var seriesResource = $resource("/admin/series/:id/:action", {id: "@id"}, {
		update: {
			method: "Post",
			params: {action: "edit"}
		}
	});

	return {
		post: postResource,
		page: pageResource,
		uploads: uploadResource,
		series: seriesResource
	}
}])