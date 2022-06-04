import { Router } from 'express';
import usersRoute from './routes/users.route.js';
const router = Router();

// user routes
router.use('/users', usersRoute);

export default router;
