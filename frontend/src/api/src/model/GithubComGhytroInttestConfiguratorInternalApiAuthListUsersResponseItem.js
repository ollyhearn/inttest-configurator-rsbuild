/**
 * IntTest configurator
 * idk what to write here it's just a swagger
 *
 * The version of the OpenAPI document: 2.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem model module.
 * @module model/GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem
 * @version 2.0
 */
class GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem {
    /**
     * Constructs a new <code>GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem</code>.
     * @alias module:model/GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem
     */
    constructor() { 
        
        GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem} obj Optional instance to populate.
     * @return {module:model/GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem} The populated <code>GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem();

            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'String');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('roles')) {
                obj['roles'] = ApiClient.convertToType(data['roles'], ['String']);
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} created_at
 */
GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem.prototype['created_at'] = undefined;

/**
 * @member {Number} id
 */
GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem.prototype['id'] = undefined;

/**
 * @member {Array.<String>} roles
 */
GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem.prototype['roles'] = undefined;

/**
 * @member {String} username
 */
GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem.prototype['username'] = undefined;






export default GithubComGhytroInttestConfiguratorInternalApiAuthListUsersResponseItem;

