<template>
<div class="login">
    <div class="background-image"></div>
    <form class="login-form">
        <h2>CryptoBank</h2>
        <img class="icon" src="https://cdn2.iconfinder.com/data/icons/security-1/512/electronic_lock-512.png" alt="CryptoBank Icon">
            <md-input-container>
                <label>Username</label>
                <md-input v-model="initialValue"></md-input>
            </md-input-container>
            <md-input-container md-has-password>
    <label>Password</label>
    <md-input type="password"></md-input>
  </md-input-container>
            <md-button class="md-raised md-primary">Anmelden</md-button>
    </form>    
</div>
  
</template>

<script>
export default {
  name: 'Login',
};
</script>

<style>
.login {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.background-image {
  background-image: url("http://img04.deviantart.net/e9ce/i/2014/013/4/2/digital_background___plexus_by_gentaro_de_averial-d721rui.png");
  background-size: cover;
  display: block;
  height: 800px;
  filter: blur(5px);
  -webkit-filter: blur(5px);
  position: fixed;
  left: 0px;
  right: 0px;
  top: 0px;
  z-index: 1;
}

.login-form {
  z-index: 2;
  color: black;
  max-width: 500px;
  margin-top: 50px;
  background-color: rgba(255, 255, 255, 0.80);
  width: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  border-radius: 3px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.25);
}

.icon {
    width: 6em;
}
</style>


