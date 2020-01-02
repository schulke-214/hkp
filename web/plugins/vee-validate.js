import Vue from 'vue';
import { ValidationProvider, ValidationObserver, extend, localize } from 'vee-validate';
import * as rules from 'vee-validate/dist/rules';
import de from 'vee-validate/dist/locale/de';

Object.keys(rules).map(rule => {
	extend(rule, {
		...rules[rule],
		message: de.messages[rule]
	});
});

extend('int32', {
	validate: val => -2100000000 < val && val < 2100000000,
	message: (_, { _value_ }) =>
		`Der eingegebene Wert kann nicht verarbeitet werden, da er zu ${_value_ > 0 ? 'groß' : 'klein'} ist`
});

localize('de', {
	messages: {
		max: (field, { length }) => `${field} darf maximal ${length} Zeichen beinhalten`,
		min_value: (field, { min }) => `${field} muss größer als ${min - 1} sein`,
		max_value: (field, { max }) => `${field} darf maximal ${max} sein`
	}
});

Vue.component('validation-provider', ValidationProvider);
Vue.component('validation-observer', ValidationObserver);
