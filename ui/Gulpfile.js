const gulp = require('gulp');
const pug = require('gulp-pug');
const sass = require('gulp-sass')(require('sass'));
const autoprefixer = require('gulp-autoprefixer');
const browsersync = require('browser-sync').create();
const ts = require('gulp-typescript');
const tsProject = ts.createProject('tsconfig.json');

// BrowserSync for live reload on file saves
function browserSync(done) {
  browsersync.init({
    server: {
      baseDir: ['./'],
      directory: true,
      stream: true
    },
    // Specifiy a socket implementation of your website
    //socket: {

       // For using a socket implementation of your website,
       // configure your DNS records to point to the ip address of yourhost machine.
       // Specify the socket name next to the field 'namespace' equalto your chosen DNS record pointing
       // to you host machine. ex: 'your.dns.record'

       // namespace: "gulp.mintymint.info"
    //},
    // disable scroll, click, page refresh etc... to be disabled across all clients
    // ghostMode: true,
    port: 8080
  });
  done();
}
// BrowserSync Reload

// Login
function browserSyncReload(done) {
  browsersync.reload();
  done();
}
function pugPrologin() {
   return (
      gulp.src('views/loginpage/pug/login.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('views'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};
function sassPrologin() {
   return (
      gulp.src('views/loginpage/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('assets/loginpage'))
      .pipe(browsersync.reload({stream: true}))
   );
};
function tsPrologin(){
   var tsResult = gulp.src('views/loginpage/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());

   return tsResult.js.pipe(gulp.dest('assets/loginpage'));
}
// Login

// signup
function pugProSignup() {
   return (
      gulp.src('views/signup/pug/signup.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('views'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};
function sassProSignup() {
   return (
      gulp.src('views/signup/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('assets/signup'))
      .pipe(browsersync.reload({stream: true}))
   );
};
function tsProSignup(){
   var tsResult = gulp.src('views/signup/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());

   return tsResult.js.pipe(gulp.dest('assets/signup'));
}
// signup

// Errorpage
function pugProErrorpage() {
   return (
      gulp.src('views/errorpage/pug/error.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('views'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};

function sassProErrorpage() {
   return (
      gulp.src('views/errorpage/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('assets/errorpage'))
      .pipe(browsersync.reload({stream: true}))
   );
};

function tsProErrorpage(){
   var tsResult = gulp.src('views/errorpage/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());
   return tsResult.js.pipe(gulp.dest('assets/errorpage'));
};
// Errorpage
// home
function pugProHome() {
   return (
      gulp.src('views/home/pug/home.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('views'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};

function sassProHome() {
   return (
      gulp.src('views/home/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('assets/home'))
      .pipe(browsersync.reload({stream: true}))
   );
};

function tsProHome(){
   var tsResult = gulp.src('views/home/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());
   return tsResult.js.pipe(gulp.dest('assets/home'));
};
// home

// Portal
function pugProportal() {
   return (
      gulp.src('views/portal/pug/portal.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('views'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};
function sassProportal() {
   return (
      gulp.src('views/portal/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('assets/portal'))
      .pipe(browsersync.reload({stream: true}))
   );
};
function tsProportal(){
   var tsResult = gulp.src('views/portal/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());

   return tsResult.js.pipe(gulp.dest('assets/portal'));
}
// Portal

// Watch files
function watchFiles() {

   //gulp.watch(['src/css/**/*.sass','src/pug/**/*.pug','src/ts/**/*.ts'], gulp.series([sassPro,pugPro,tsPro]));

   gulp.watch('views/loginpage/sass/**/*.sass',sassPrologin);
   gulp.watch('views/loginpage/pug/**/*.pug',pugPrologin);
   gulp.watch('views/loginpage/ts/**/*.ts',tsPrologin);

   gulp.watch('views/portal/sass/**/*.sass',sassProportal);
   gulp.watch('views/portal/pug/**/*.pug',pugProportal);
   gulp.watch('views/portal/ts/**/*.ts',tsProportal);

   gulp.watch('views/errorpage/sass/**/*.sass',sassProErrorpage);
   gulp.watch('views/errorpage/pug/**/*.pug',pugProErrorpage);
   gulp.watch('views/errorpage/ts/**/*.ts',tsProErrorpage);
   
   gulp.watch('views/home/sass/**/*.sass',sassProHome);
   gulp.watch('views/home/pug/**/*.pug',pugProHome);
   gulp.watch('views/home/ts/**/*.ts',tsProHome);
   
   gulp.watch('views/signup/sass/**/*.sass',sassProSignup);
   gulp.watch('views/signup/pug/**/*.pug',pugProSignup);
   gulp.watch('views/signup/ts/**/*.ts',tsProSignup);
   
   gulp.watch([
      'assets/loginpage/index.html',
      'assets/loginpage/style.css',
      'assets/loginpage/index.js',

      'assets/portal/index.html',
      'assets/portal/style.css',
      'assets/portal/index.js',

      'assets/errorpage/index.html',
      'assets/errorpage/style.css',
      'assets/errorpage/index.js',

      'assets/home/index.html',
      'assets/home/style.css',
      'assets/home/index.js',

      'assets/signup/index.html',
      'assets/signup/style.css',
      'assets/signup/index.js'
   ], 
   gulp.series(browserSyncReload));
};
// Watch files

// define tasks to process
const build = gulp.series(
   pugPrologin,
   pugProportal,
   pugProErrorpage,
   pugProHome,
   pugProSignup,

   sassPrologin,
   sassProportal,
   sassProErrorpage,
   sassProHome,
   sassProSignup,

   tsPrologin,
   tsProportal,
   tsProErrorpage,
   tsProHome,
   tsProSignup,

   browserSync, 
   watchFiles
);
// define tasks to process

// export tasks
exports.default = build;