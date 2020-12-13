import About from "../../build/rendered_views/About.js";
import SuccessBackendMock from "../../app/SuccessBackendMock.js";

const assert = require('assert');

describe('AboutView', function() {
  it('should contain text', async function () {
    const view = new About(new SuccessBackendMock(), null, {});
    const text = await view.getHtml();
    assert(text.includes("some text"));
  });
});


