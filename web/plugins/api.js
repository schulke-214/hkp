import { APIClient } from '~/api';

export default ({ app }, inject) => {
	inject('api', new APIClient(app.$axios));
};
