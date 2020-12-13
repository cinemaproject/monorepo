/**
 *  Model Films
 * @param {*} sequelize
 * @param {*} DataTypes
 */
const Films = (sequelize, DataTypes) => {
  const Films = sequelize.define(
    "films", // name of Model
    {
      // fields
      title: { type: DataTypes.STRING, notEmpty: false },
      type: { type: DataTypes.STRING, notEmpty: true },
      start_year: { type: DataTypes.STRING, notEmpty: true },
      end_year: { type: DataTypes.STRING, notEmpty: true },
      runtime_minutes: { type: DataTypes.STRING, notEmpty: true }
      // description: { type: DataTypes.TEXT, is: ["^[a-z]{10,}$", "i"] },
    }, {
        timestamps: false
    }
  );
  return Films;
};

module.exports = Films;
