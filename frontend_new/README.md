# Get started with local development

**This guide does not use kubernetes**

1. Install Node.js: https://nodejs.org/en/
2. Install npm (if it's not installed with nodejs): https://www.npmjs.com/get-npm
3. Follow these steps iteratively:
```sh
cd monorepo/frontend_new
npm install # Only if new modules have been added to package.json
npx gulp serve
```

4. Running tests:
```sh
cd monorepo/frontend_new
npm test
```
