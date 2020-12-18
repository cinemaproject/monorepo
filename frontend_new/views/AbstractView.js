export default class AbstractView {
    constructor(backendService, params, document) {
        this.backendService = backendService;
        this.params = params;
        this.document = document;
    }

    setTitle(title) {
        this.document.title = title;
    }

    async getHtml() {
        return "";
    }
}
