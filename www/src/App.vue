<template>
  <v-app id="keep">
    <v-navigation-drawer
            v-model="drawer"
            fixed
            clipped
            class="grey lighten-4"
            app>
      <v-list dense
              class="grey lighten-4">
        <template v-for="(item, i) in items">
          <v-layout
                  v-if="item.heading"
                  :key="i"
                  row
                  align-center>
            <v-flex xs6>
              <v-subheader v-if="item.heading">
                {{ item.heading }}
              </v-subheader>
            </v-flex>
          </v-layout>

          <v-divider
                  v-else-if="item.divider"
                  :key="i"
                  dark
                  class="my-3"
          ></v-divider>
          <v-list-tile
                  v-else
                  :key="i"
                  :to="item.link">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title class="grey--text">
                {{ item.text }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar color="amber" app absolute clipped-left>
      <v-toolbar-side-icon @click="drawer = !drawer"></v-toolbar-side-icon>

          <v-img
                  :src="require('./assets/logo-plamienok.png')"
                  contain
                  max-height="55"
                  max-width="140"
                  @click="dashboard()"
          ></v-img>

      <span class="title ml-3 mr-5">Plamienok&nbsp;<span class="font-weight-light">App</span></span>
      <v-text-field
              solo-inverted
              flat
              hide-details
              label="Hladať"
              prepend-inner-icon="search"
      ></v-text-field>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>

    export default {
        name: 'App',

        components: {
        },

        data: () => ({
            drawer: null,
            items: [
                { heading: 'Home care' },
                { icon: 'lightbulb_outline  ', text: 'Návštevy' },
                { icon: 'touch_app', text: 'Telefón' },
                { divider: true },
                { heading: 'Poradňa' },
                { icon: 'add', text: 'Stretnutia' },
                { divider: true },
                { heading: 'Štatistiky' },
                { icon: 'settings', text: 'Hodiny' },
                { icon: 'chat_bubble', text: 'Autá' },
                { heading: 'Administrácia' },
                { icon: 'phonelink', text: 'Užívatelia' },
                { icon: 'keyboard', text: 'Autá' },
                { icon: 'phonelink', text: 'Exporty' },
            ]
        }),

        methods: {
          dashboard: function() {
            return this.$router.push('/');
          }
        },

        props: {
            source: String
        }


    }
</script>

<style lang="stylus">
  #keep main .container {
    height: 660px;
  }

  .navigation-drawer__border {
    display: none;
  }

  .text {
    font-weight: 400;
  }
</style>

