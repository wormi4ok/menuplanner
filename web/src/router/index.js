import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '@/views/Home.vue';
import Recipes from '@/views/Recipes.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/recipes',
    name: 'Recipes',
    component: Recipes,
  },
];

const router = new VueRouter({
  routes,
});

export default router;
