import React, { useState } from "react";

import "./UserEditor.css";
import { Button, Modal, Table, message } from "antd";
import { EditOutlined } from "@ant-design/icons";
import UsersApi from "../../api/src/api/UsersApi";

const UserEditPage = () => {
  const userApi = new UsersApi();
  const roleById = fetchRoleIdCache(userApi);
  const permById = fetchPermIdCache(userApi);
  const userData = fetchUserList(userApi);
  const [isUserEditOpen, setUserEditOpen] = useState(false);

  const userListTableCols = [
    {
      title: "Логин",
      dataIndex: "username",
    },
    {
      title: "Роли",
      dataIndex: "roleNames",
      render: (_, roleIds) => (
        <span>
          {roleIds
            .map((roleId) => {
              return roleById[roleId].name;
            })
            .join(", ")}
        </span>
      ),
    },
    {
      title: "Редактирование",
      render: (_, record) => (
        <Button
          shape="circle"
          icon={<EditOutlined />}
          onclick={setUserEditOpen(true)}
          size="large"
        />
      ),
    },
  ];

  const userListTableDataSource = userData.map((user) => {
    return {
      username: user.username,
      roleNames: user.roles
        .map((roleId) => {
          return roleById[roleId].name;
        })
        .join(", "),
    };
  });

  return (
    <>
      <Table
        columns={userListTableCols}
        dataSource={userListTableDataSource}
      ></Table>
      <Modal
        title="Редактирование пользователя"
        open={isUserEditOpen}
        on
      ></Modal>
    </>
  );
};

function fetchRoleIdCache(userApi) {
  let result = undefined;
  userApi.listRoles((e, data, response) => {
    if (e) {
      message.error(response?.body?.message);
      return;
    }
    result = {};
    for (const r of response) {
      const role = r?.body;
      result[role?.id] = role;
    }
  });
  return result;
}

function fetchPermIdCache(userApi) {
  let result = undefined;
  userApi.listPerms((e, data, response) => {
    if (e) {
      message.error(response?.body?.message);
      return;
    }
    for (const r of response) {
      const perm = r?.body;
      result[perm?.id] = perm;
    }
    return result;
  });
  return result;
}

function fetchUserList(userApi) {
  let result = undefined;
  userApi.listUsers((e, data, response) => {
    if (e) {
      message.error(response?.body?.message);
      return;
    }
    result = response?.body;
  });
  return result;
}

function submitUpdateUserData(userApi) {
  return (formData) => {};
}

export default UserEditPage;
