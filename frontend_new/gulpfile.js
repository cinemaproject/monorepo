const { series, src, dest, parallel } = require('gulp');
const injectTemplate = require('gulp-inject-template');
const concatCss = require('gulp-concat-css');
const copy = require('gulp-copy');
const browserSync = require('browser-sync').create();

function render_templates() {
  return src('views/**/*.js', { buffer: false })
    .pipe(injectTemplate({variable: 'parameters'}))
    .pipe(dest('./build/rendered_views/'));
}

function compile_css() {
  return src('styles/**/*.css')
    .pipe(concatCss("main.css"))
    .pipe(dest('./build/styles'));
}

function deployViewJS() {
  return src(['./build/rendered_views/*.js']).pipe(copy('./dist/static/js/views/', { prefix: 2 }));
}

function deployCSS() {
  return src('./build/styles/main.css').pipe(copy('./dist/static/css/', { prefix: 2 }));
}

function deployJS() {
  return src('./app/main.js').pipe(copy('./dist/static/js/', { prefix: 1 }));
}

function deployHTML() {
  return src('./index.html').pipe(copy('./dist/'));
}

function deployImages() {
  return src(['./images/*.png', './images/*.jpg']).pipe(copy('./dist/static'));
}

function serve(cb) {
  browserSync.init({
    server: {
      baseDir: "./dist"
    }
  });
}

const cssPipeline = series(compile_css);
const jsPipeline = series(render_templates);
const deploy = series(deployCSS, deployJS, deployViewJS, deployHTML, deployImages);

const defaultPipeline = series(
  parallel(cssPipeline, jsPipeline),
  deploy
);

module.exports.default = defaultPipeline;
module.exports.serve = series(
  defaultPipeline,
  serve
); 
