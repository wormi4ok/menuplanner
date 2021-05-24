import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '@/views/Home.vue';
import Recipes from '@/views/Recipes.vue';
import Login from '@/views/Login.vue';
import middleware from './middleware';

Vue.use(VueRouter);

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

const router = new VueRouter({
  routes,
  mode: 'history',
});

router.beforeEach(middleware.initUser);
router.beforeEach(middleware.checkAccess);

export default router;
