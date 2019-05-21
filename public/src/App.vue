<template>
  <div id="app">
    <Login />
    <a-layout id="components-layout-demo-top" class="layout">
      <a-layout-header>
        <div class="logo">
          LINKR
        </div>
        <a-menu
          theme="dark"
          mode="horizontal"
          :defaultSelectedKeys="['nav:dashboard']"
          :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="nav:dashboard">
            <router-link to="/dashboard">Dashboard</router-link>
          </a-menu-item>
          <a-menu-item key="nav:settings" to="/settings">
            <router-link to="/settings">Settings</router-link>
          </a-menu-item>
        </a-menu>
      </a-layout-header>
      <a-layout-content style="padding: 0 50px">
        <a-breadcrumb style="margin: 16px 0">
          <a-breadcrumb-item>Home</a-breadcrumb-item>
          <a-breadcrumb-item>List</a-breadcrumb-item>
          <a-breadcrumb-item>App</a-breadcrumb-item>
        </a-breadcrumb>
        <div :style="{ background: '#fff', padding: '24px', minHeight: '280px' }">
          <router-view/>
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Ant Design ©2018 Created by Ant UED
      </a-layout-footer>
    </a-layout>
  </div>
</template>

<script>
  import Login from '@/components/Login'
  export default {
    components: {
      Login
    }
  }
</script>

<style>
#components-layout-demo-top .logo {
  width: 120px;
  height: 31px;
  line-height: 31px;
  margin: 16px 24px 16px 0;
  float: left;
  text-align: center;
  color: #cccccc;
  letter-spacing: 1px;
  font-weight: bold;
  cursor: default;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
 }
</style>
