import Vue from 'vue';
import Router from 'vue-router';
import Overview from '@/views/Overview';
import Login from '@/views/Login';
import 'vue-material/dist/vue-material.css';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/Login',
      name: 'Login',
      component: Login,
    },
    {
      path: '/Overview',
      name: 'Overview',
      component: Overview,
    },
  ],
});
