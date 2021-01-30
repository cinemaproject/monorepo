import AbstractView from "./AbstractView.js";

export default class SearchResults extends AbstractView {
  constructor(backendService, params, document) {
    super(backendService, params, document);
    this.setTitle("Search Results");
  }

  async getHtml() {
    const header = require("../../templates/search_results_header.html");
    const sectionHeader = require("../../templates/search_results_section_header.html");
    const sectionFooter = require("../../templates/search_results_section_footer.html");
    const cardTemplate = require("../../templates/card.html");

    let cards = "";

    const films = await this.backendService.searchFilms(this.params.query);

    if (films != null && films.length > 0) {
      cards += sectionHeader({ title: "Films" });
      films.forEach(element => {
        cards += cardTemplate({
          title: element.title,
          image_url: element.poster_url,
          desc: "Lorem impsum dolor sit amet"
        }) + "\n";
      });
      cards += sectionFooter();
    }

    const people = await this.backendService.searchPeople(this.params.query);

    if (people != null && people.length > 0) {
      cards += sectionHeader({ title: "People" });
      people.forEach(element => {
        cards += cardTemplate({
          title: element.primary_name,
          image_url: element.photo_url,
          desc: "Short actor bio"
        });
      });
      cards += sectionFooter();
    }
    return header() + cards;
  }
}
