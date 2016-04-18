window.Mock = require("lib/mock.js");


// Mock.mockjax(jQuery);
Mock.setup({
    timeout: "50-300"
})

// 设备列表
Mock.mock('api/devices', {
    code: 0,
    data: [
        {
            id: "1",
            name: "sdf",
            url: "tes.com",
            type: 1,
            enabled: false
        },
        {
            id: "2",
            name: "sdf",
            url: "tes.com",
            type: 2,
            enabled: true
        }
    ]
})

// 设备url解析，预览
Mock.mock(/api\/device\/url\/parse\?.*/, {
    code: 0,
    data: "/public/img/DeepGlint.jpg"
})

// 设备url解析，预览
Mock.mock("api/tasks", {
    code: 0,
    data: [
        {
            id: "1",
            name: "test",
            url: "sdf.com",
            type: 1,
            state: 1,
            create_time: 1456821662265,
            execute_time: 1456821662265
        },
        {
            id: "2",
            name: "test",
            url: "sdf.com",
            type: 2,
            state: 2,
            create_time: 1456821662265,
            execute_time: 1456821662265
        }
    ]
})

