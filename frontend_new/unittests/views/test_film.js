import Film from "../../build/rendered_views/Film.js";
import SuccessBackendMock from "../../app/SuccessBackendMock.js";

const assert = require('assert');

describe('FilmView', function() {
  it('should contain film id', async function () {
    const view = new Film(new SuccessBackendMock(), {"id": "1"}, {});
    const text = await view.getHtml();
    assert(text.includes("1"));
  });
});


