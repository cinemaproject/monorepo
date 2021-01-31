import SearchResults from "../../build/rendered_views/SearchResults.js";
import SuccessBackendMock from "../../app/SuccessBackendMock.js";

const assert = require('assert');

describe('SearchResultsView', function() {
  it('should contain cards', async function () {
    const view = new SearchResults(new SuccessBackendMock(), {"query": "anything"}, {});
    const text = await view.getHtml();
    assert(text.includes("Results"));
    assert(text.includes("Films"));
    assert(text.includes("Sample Film"));
    assert(text.includes("People"));
    assert(text.includes("John Doe"));
  });
});
