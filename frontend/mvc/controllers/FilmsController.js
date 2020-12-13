const db = require('../models/index.js');
/**
 * Class Films Controller
 */
class FilmsController {
  /**
   * Liste of Films
   * @param {*} req
   * @param {*} res
   */
  liste(req, res) {
    db.films.findAndCountAll({
      // where: {
      //   title: {
      //     [Op.like]: 'foo%'
      //   }
      // },
      // offset: 10,
      limit: 10
    }).then(films =>
      res.render("films/index", { films })
    );
  }

  /**
   * Show a article
   * @param {*} req
   * @param {*} res
   */
  show(req, res) {
    const id = req.params.id;
    db.Films.findById(id).then(article =>
      res.render("films/show", { article })
    );
  }
}

module.exports = FilmsController;