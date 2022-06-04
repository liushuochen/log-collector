import { Router } from 'express';
import usersController from '../controllers/users.controller.js';
const usersRouter = Router();

// GET /users/downloadlist
usersRouter.get('/downloadlist', (req, res) => {
  usersController.getDownloadList(req, res);
});

export default usersRouter;
