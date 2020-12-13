const { series, src, dest, parallel } = require('gulp');
const injectTemplate = require('gulp-inject-template');
const concatCss = require('gulp-concat-css');
const copy = require('gulp-copy');
const browserSync = require('browser-sync').create();

function render_templates(cb) {
  src('views/**/*.js', { buffer: false })
    .pipe(injectTemplate({variable: 'parameters'}))
    .pipe(dest('./build/rendered_views/'));
  cb();
}

function compile_css(cb) {
  src('styles/**/*.css')
    .pipe(concatCss("main.css"))
    .pipe(dest('./build/styles'));
  cb();
}

function deploy(cb) {
  src(['./build/rendered_views/*.js']).pipe(copy('./dist/static/js/views/', { prefix: 2 }));
  src('./styles/main.css').pipe(copy('./dist/static/css/', { prefix: 1 }));
  src('./app/main.js').pipe(copy('./dist/static/js/', { prefix: 1 }));
  src('./index.html').pipe(copy('./dist/'));
  cb();
}

function serve(cb) {
  browserSync.init({
    server: {
      baseDir: "./dist"
    }
  });
  //cb();
}

exports.default = series(
  parallel(render_templates, compile_css),
  deploy,
  serve
);
