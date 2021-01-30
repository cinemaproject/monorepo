import AbstractView from "./AbstractView.js";

export default class SearchResults extends AbstractView {
  constructor(backendService, params, document) {
    super(backendService, params, document);
    this.setTitle("Search Results");
  }

  async getHtml() {
    const header = require("../../templates/search_results_header.html")
    const cardTemplate = require("../../templates/film_card.html")

    let cards = ""

    const films = await this.backendService.searchFilms(this.params.query);

    films.forEach(element => {
      cards += cardTemplate({
        film: element
      }) + "\n";
    });

    return header() + cards;
  }
}
