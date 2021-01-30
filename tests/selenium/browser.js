const webdriver = require('selenium-webdriver');

module.exports.createBrowser = async () => {
  let browser = process.env.BROWSER || "chrome";
  let capabilities;
  switch (browser) {
    case "chrome": {
      require("chromedriver");
      capabilities = webdriver.Capabilities.chrome();
      var args = [
        "--headless",
        "--disable-gpu",
        "--disable-sndbox"
      ];

      capabilities.set("chromeOptions", {
        args: args
      });
      break;
    }
    case "safari": {
      capabilities = webdriver.Capabilities.safari();
      break;
    }
    case "firefox": {
      require("geckodriver");
      capabilities = webdriver.Capabilities.firefox();
      break;
    }
  }

  let instance;

  try {
    instance = await new webdriver.Builder()
      .withCapabilities(capabilities)
      .build();
  } catch (err) {
    console.error(err);
  }

  return instance;
};
module.exports.base_url = process.env.BASE_URL || "http://localhost:3000";
