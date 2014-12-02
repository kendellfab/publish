pubApp.controller("EditController", ["$scope", "res", function($scope, res) {
	$scope.post = backend.post;


	$scope.savePost = function() {
		// alert(JSON.stringify($scope.post));
		res.post.update($scope.post);
	};
}]);