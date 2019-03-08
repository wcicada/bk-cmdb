/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package initial

import (
	"configcenter/src/auth/authcenter"
)

var expectBizResourceType = []authcenter.ResourceType{
	{
		ResourceTypeID:       authcenter.BizModelGroup,
		ResourceTypeName:     "模型分组",
		ParentResourceTypeID: "",
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizModel,
		ResourceTypeName:     "模型",
		ParentResourceTypeID: authcenter.BizModelGroup,
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizInstance,
		ResourceTypeName:     "实例",
		ParentResourceTypeID: authcenter.BizModel,
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "查询",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizCustomQuery,
		ResourceTypeName:     "动态分组",
		ParentResourceTypeID: "",
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "查询",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizProcessInstance,
		ResourceTypeName:     "进程",
		ParentResourceTypeID: "",
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "查询",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.BindModule,
				ActionName:        "绑定到模块",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizTopoInstance,
		ResourceTypeName:     "拓扑",
		ParentResourceTypeID: "",
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Create,
				ActionName:        "新建",
				IsRelatedResource: false,
			},
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "删除",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "查询",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.HostTransfer,
				ActionName:        "主机转移",
				IsRelatedResource: true,
			},
		},
	},
	{
		ResourceTypeID:       authcenter.BizHostInstance,
		ResourceTypeName:     "主机",
		ParentResourceTypeID: "",
		Actions: []authcenter.Action{
			{
				ActionID:          authcenter.Edit,
				ActionName:        "编辑",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.Delete,
				ActionName:        "查询",
				IsRelatedResource: true,
			},
			{
				ActionID:          authcenter.ModuleTransfer,
				ActionName:        "转资源池",
				IsRelatedResource: true,
			},
		},
	},
}
