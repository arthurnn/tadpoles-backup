const fs = require('fs');
const fse = require('fs-extra');
const bookmarklet = require('bookmarklet');
const Mustache = require('mustache');


const scriptSource = './reset_bookmarklet/reset_script.js';
const templateSource = './reset_bookmarklet/bookmarklet_info_template.mustache';

async function build() {
    const scriptRaw = fs.readFileSync(scriptSource, {encoding: 'utf8'});
    const templateRaw = fs.readFileSync(templateSource).toString();

    const data = bookmarklet.parseFile(scriptRaw);

    if (data.errors) {
        const msg = data.errors.join('\n');
        console.error(`[ERROR] bookmarklet: ${msg}`);
        process.exit(1);
    }

    const templateArgs = {
        name: 'reset-tadpoles-password',
        minified: await bookmarklet.convert(data.code, data.options),
    };

    const rendered = Mustache.render(templateRaw, templateArgs);

    const build_dir = './dist';

    fse.ensureDirSync(build_dir);

    fs.writeFileSync(`${build_dir}/reset-tadpoles-password.html`, rendered);
}

build()
    .catch(e => console.log('build failure', e))
    .then(() => console.log('build complete'));
