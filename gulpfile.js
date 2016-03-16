var gulp = require('gulp'),
  $ = require('gulp-load-plugins')(),
  source = require('vinyl-source-stream'),
  del = require('del');

var project = $.typescript.createProject('tsconfig.json', {typescript: require('typescript')});

gulp.task('compile', function () {
  return gulp.src("src/**/*.{ts,tsx}")
    .pipe($.sourcemaps.init())
    .pipe($.typescript(project))
    .js
    .pipe($.sourcemaps.write())
    .pipe(gulp.dest('.'));
});

gulp.task('bundle', ['compile'], function () {
  return gulp.src(['.tmp/tsBundle.js', 'src/**/*.js'])
    .pipe($.browserify({
      debug: true
    }))
    .pipe($.rename('bundle.js'))
    .pipe(gulp.dest('dist'));
});

gulp.task('clean', function (done) {
  del(['.tmp'], done.bind(this));
});

gulp.task('watch', function () {
  $.watch('src/**/*.{ts,tsx}', ['bundle']);
});

gulp.task('default', ['watch']);
