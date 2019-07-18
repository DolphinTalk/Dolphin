syntax = "proto3";
package pb;
import "google/api/annotations.proto";

service Dolphin {
	rpc registOnline() returns () { }

	//将消息推送到端
	rpc pushMsg() returns() { }

	//将消息发送给Dolphin服务器
	//如果不在本机，帮忙路由
	rpc sendMsg() returns () {}

	//
	rpc getFriendList returns() {}
}


