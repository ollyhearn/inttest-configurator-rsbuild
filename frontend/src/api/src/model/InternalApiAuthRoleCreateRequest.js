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
 * The InternalApiAuthRoleCreateRequest model module.
 * @module model/InternalApiAuthRoleCreateRequest
 * @version 2.0
 */
class InternalApiAuthRoleCreateRequest {
    /**
     * Constructs a new <code>InternalApiAuthRoleCreateRequest</code>.
     * @alias module:model/InternalApiAuthRoleCreateRequest
     */
    constructor() { 
        
        InternalApiAuthRoleCreateRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InternalApiAuthRoleCreateRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InternalApiAuthRoleCreateRequest} obj Optional instance to populate.
     * @return {module:model/InternalApiAuthRoleCreateRequest} The populated <code>InternalApiAuthRoleCreateRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InternalApiAuthRoleCreateRequest();

            if (data.hasOwnProperty('desc')) {
                obj['desc'] = ApiClient.convertToType(data['desc'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('perm_ids')) {
                obj['perm_ids'] = ApiClient.convertToType(data['perm_ids'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} desc
 */
InternalApiAuthRoleCreateRequest.prototype['desc'] = undefined;

/**
 * @member {String} name
 */
InternalApiAuthRoleCreateRequest.prototype['name'] = undefined;

/**
 * @member {Array.<Number>} perm_ids
 */
InternalApiAuthRoleCreateRequest.prototype['perm_ids'] = undefined;






export default InternalApiAuthRoleCreateRequest;
