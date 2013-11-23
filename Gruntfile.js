module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    less: {
      development: {
        options: {
          paths: ["assets/css"]
        },
        files: {
          "assets/css/dev/main.css": "assets/css/main.less"
        }
      },
      production: {
        options: {
          paths: ["assets/css"],
          cleancss: true
        },
        files: {
          "static/main.css": "assets/css/main.less"
        }
      }
    },
    watch: {
      options: {
        livereload: true
      },
      less: {
        files: ["assets/css/main.less"],
        tasks: ["less"],
      }
    },
  });

  // Load the plugin that provides the "uglify" task.
  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['less']);

};