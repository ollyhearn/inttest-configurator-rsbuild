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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.IntTestConfigurator);
  }
}(this, function(expect, IntTestConfigurator) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ConfiguratorInternalApiAuthCreateUserRequest', function() {
    it('should create an instance of ConfiguratorInternalApiAuthCreateUserRequest', function() {
      // uncomment below and update the code to test ConfiguratorInternalApiAuthCreateUserRequest
      //var instance = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest();
      //expect(instance).to.be.a(IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest);
    });

    it('should have the property password (base name: "password")', function() {
      // uncomment below and update the code to test the property password
      //var instance = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest();
      //expect(instance).to.be();
    });

    it('should have the property roles (base name: "roles")', function() {
      // uncomment below and update the code to test the property roles
      //var instance = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest();
      //expect(instance).to.be();
    });

    it('should have the property username (base name: "username")', function() {
      // uncomment below and update the code to test the property username
      //var instance = new IntTestConfigurator.ConfiguratorInternalApiAuthCreateUserRequest();
      //expect(instance).to.be();
    });

  });

}));
