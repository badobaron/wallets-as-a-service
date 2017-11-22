<template>
  <div>
    <header-component></header-component>
    <md-card>
      <md-card-content>
        This is the overview page
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import HeaderComponent from '@/components/Header';

export default {
  components: {
    HeaderComponent,
  },
};
</script>