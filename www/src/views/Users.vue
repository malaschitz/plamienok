<!--
  - Developed by Richard Malaschitz on 2/12/19 1:16 PM
  - Last modified 2/12/19 1:16 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-toolbar
      flat
      color="white"
    >
      <v-toolbar-title>Užívatelia</v-toolbar-title>
      <v-divider
        class="mx-2"
        inset
        vertical
      />
      <v-spacer />
      <v-checkbox
        v-model="deleted"
        label="Zrušení užívatelia"
      />
      <v-dialog
        v-model="dialog"
        max-width="500px"
      >
        <v-btn
          slot="activator"
          color="primary"
          dark
          class="mb-2"
        >
          Nový užívateľ
        </v-btn>
        <v-card>
          <v-form
            ref="userform"
            v-model="valid"
            lazy-validation
          >
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex
                    xs12
                    sm8
                    md6
                  >
                    <v-text-field
                      v-model="editedItem.Name"
                      label="Meno"
                      :rules="[rules.required]"
                      required
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm8
                    md6
                  >
                    <v-text-field
                      v-model="editedItem.Email"
                      label="Email"
                      :rules="[rules.required, rules.email]"
                      required
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-checkbox
                      v-model="editedItem.RoleAdmin"
                      label="Administrátor"
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-checkbox
                      v-model="editedItem.RoleDoctor"
                      label="Doktor"
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-checkbox
                      v-model="editedItem.RoleNurse"
                      label="Sestra"
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-checkbox
                      v-model="editedItem.RoleSoc"
                      label="Sociálny pracovník"
                    />
                  </v-flex>
                  <v-flex
                    xs12
                    sm6
                    md4
                  >
                    <v-checkbox
                      v-model="editedItem.RolePsych"
                      label="Psychológ"
                    />
                  </v-flex>
                </v-layout>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer />
              <v-btn
                color="blue darken-1"
                flat
                @click="close"
              >
                Zrušiť
              </v-btn>
              <v-btn
                color="blue darken-1"
                flat
                @click="save"
              >
                Uložiť
              </v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-dialog>
    </v-toolbar>

    <v-data-table
      :headers="headers"
      :pagination.sync="pagination"
      :items="filteredUsers"
      class="elevation-1"
      rows-per-page-text="Počet riadkov"
    >
      <template
        slot="items"
        slot-scope="props"
      >
        <td class="text-xs-left">
          {{ props.item.Name }}
        </td>
        <td class="text-xs-left">
          {{ props.item.Email }}
        </td>
        <td class="text-xs-right">
          <v-checkbox
            v-model="props.item.RoleAdmin"
            disabled
          />
        </td>
        <td class="text-xs-right">
          <v-checkbox
            v-model="props.item.RoleDoctor"
            disabled
          />
        </td>
        <td class="text-xs-right">
          <v-checkbox
            v-model="props.item.RoleNurse"
            disabled
          />
        </td>
        <td class="text-xs-right">
          <v-checkbox
            v-model="props.item.RoleSoc"
            disabled
          />
        </td>
        <td class="text-xs-right">
          <v-checkbox
            v-model="props.item.RolePsych"
            disabled
          />
        </td>
        <td class="justify-center layout px-0">
          <v-icon
            small
            class="mr-2"
            @click="editItem(props.item)"
          >
            edit
          </v-icon>
          <v-icon
            v-if="!deleted"
            small
            @click="deleteItem(props.item)"
          >
            delete
          </v-icon>
          <v-icon
            v-else
            small
            @click="deleteItem(props.item)"
          >
            autorenew
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
    export default {
        name: "Users",

        data: () => ({
          headers: [{text: 'Meno', value: 'Name', align: 'left'},{text: 'Email', value: 'Email'},
            {text: 'Admin', value: 'RoleAdmin'},{text: 'Doktor', value: 'RoleDoctor'},{text: 'Sestra', value: 'RoleNurse'},
            {text: 'Soc. prac.', value: 'RoleSoc'},{text: 'Psychológ', value: 'RolePsych'},{text: 'Akcie', sortable: false }],
          users: [],
          pagination: {rowsPerPage:10},
          editedIndex: -1,
          editedItem: {
            ID: '',
            Name: '',
            Email: '',
            RoleAdmin: false,
            RoleNurse: false,
            RoleDoctor: false,
            RoleSoc: false,
            RolePsych: false,
          },
          dialog: false,
          rules: {
            required: value => !!value || 'Required.',
            counter: value => value.length <= 20 || 'Max 20 characters',
            email: value => {
              const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
              return pattern.test(value) || 'Invalid e-mail.'
            },
          },
          valid: true,
          deleted: false,
        }),

        computed: {
          formTitle () {
            return this.editedIndex === -1 ? 'Nový užívateľ' : 'Editácia užívateľa'
          },

          filteredUsers () {
            var a = []
            for (var i=0;i<this.users.length;i++) {
              var u = this.users[i];
              if ((this.deleted && u.Deleted) || (!this.deleted && !u.Deleted)) {
                a.push(u)
              }
            }
            return a;
          },
        },

        watch: {
          dialog (val) {
            val || this.close()
          }
        },

        methods: {
          editItem: function(item) {
            this.editedIndex = this.users.indexOf(item)
            this.editedItem = Object.assign({}, item)
            this.dialog = true
          },

          deleteItem: function(item) {
            var conf;
            if (!this.deleted) {
              conf = 'Ste si istý, že chcete vymazať užívateľa ?';
            } else {
              conf = 'Ste si istý, že chcete obnoviť užívateľa ?';
            }
            if (confirm(conf)) {
              this.$axios.delete("/r/api/user/" + item.ID).then(response => {
                if (response.status == 202) {
                  this.$store.commit("alert","Chyba: " + response.data.Error);
                } else { //OK
                  this.readData();
                }
              }).catch(response => {
                console.log('WRONG',response);
                this.$store.commit("alert","Chyba: " + response);
              });
            }
          },

          close: function() {
            this.dialog = false
            setTimeout(() => {
              this.editedItem = Object.assign({}, this.defaultItem)
              this.editedIndex = -1
            }, 300)
          },

          save: function() {
            if (!this.$refs.userform.validate()) {
              return;
            }


            this.$axios.post("/r/api/user",this.editedItem).then(response => {
              if (response.status == 202) {
                this.$store.commit("alert","Chyba: " + response.data.Error);
              } else { //OK
                this.editedItem = response.data;

                if (this.editedIndex > -1) {
                  Object.assign(this.users[this.editedIndex], this.editedItem)
                } else {
                  this.users.push(this.editedItem)
                }

                this.close()
              }
            }).catch(response => {
              console.log('WRONG',response);
              this.$store.commit("alert","Chyba: " + response);
            });



          },

          readData: function() {
            this.$axios.get('/r/api/users').then(response => {
              console.log('OK',response);
              if (response.status == 202) {
                this.$store.commit("alert","Chyba: " + response.data.Error);
              } else {
                this.users = response.data;
              }
            }).catch(response => {
              console.log('WRONG',response);
              this.$store.commit("alert","Chyba: " + response);
            });
          },
        },

        mounted: function() {
          console.log("mounted users");
          this.readData();
        },
    }
</script>

<style scoped>

</style>