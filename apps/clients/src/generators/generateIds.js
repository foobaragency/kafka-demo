const { v4: uuidV4 } = require("uuid");

function generateIds() {
  const quantity = Math.floor(Math.random() * 10);

  const ids = [uuidV4(), uuidV4(), uuidV4()];

  for (let i = 0; i < quantity; i++) {
    ids.push(uuidV4());
  }

  return ids;
}

module.exports = { generateIds };
