const express = require("express");
const matrixController = require("../controllers/matrixController");
const { authenticateJWT } = require("../middleware/authMiddleware");

const router = express.Router();

router.post("/statistics", authenticateJWT, matrixController.computeStatistics);

module.exports = router;