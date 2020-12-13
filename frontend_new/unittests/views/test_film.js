import Film from "../../build/rendered_views/Film.js";
import SuccessBackendMock from "../utils/SuccessBackendMock.js";

const assert = require('assert');

describe('FilmView', function() {
  it('should contain film id', async function () {
    const view = new Film(new SuccessBackendMock(), {"id": "t00000001"}, {});
    const text = await view.getHtml();
    assert(text.includes("t000000001"));
  });
});


