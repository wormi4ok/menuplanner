import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/Home.vue';
import Recipes from '@/views/Recipes.vue';
import Login from '@/views/Login.vue';
import middleware from './middleware';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { isAuth: true },
  },
  {
    path: '/recipes',
    name: 'Recipes',
    component: Recipes,
    meta: { isAuth: true },
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(),
});

router.beforeEach(middleware.initUser);
router.beforeEach(middleware.checkAccess);

export default router;
