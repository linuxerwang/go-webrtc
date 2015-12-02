/*
Package webrtc is a golang wrapper for libwebrtc.

To provide an easier experience for users of this package, there are differences
inherent in the interface written here and the original native code WebRTC. This
allows users to use WebRTC in a more idiomatic golang way. For example, callback
mechanism has a layer of indirection that allows goroutines instead.

There is also a complication in building the dependent static library for this
to work. More later...

TODO(keroserene): More package documentation, and documentation in general.
*/
package webrtc

/*
// #cgo CPPFLAGS: -I/usr/include/c++/5.2.0 -I/usr/include/c++/5.2.0/x86_64-unknown-linux-gnu
#cgo CPPFLAGS: -Ithird_party/libwebrtc/
#cgo CXXFLAGS: -std=gnu++11 -Wno-c++0x-extensions
#cgo LDFLAGS: -Wl,-z,now -Wl,-z,relro -Wl,--fatal-warnings -Wl,-z,defs -pthread
#cgo LDFLAGS: -Wl,-z,noexecstack -fPIC -fuse-ld=gold
#cgo LDFLAGS: -B/home/serene/code/webrtc-check/src/third_party/binutils/Linux_x64/Release/bin
#cgo LDFLAGS: -Wl,--disable-new-dtags -pthread -m64 -Wl,--detect-odr-violations
// #cgo LDFLAGS: -Wl,--icf=all -Wl,-O1 -Wl,--as-needed -Wl,--gc-sections
#cgo LDFLAGS: -L/home/serene/code/go/src/github.com/keroserene/webrtc/lib
// #cgo LDFLAGS: -Llib
// libwebrtc_magic is a custom built archive based on many other ninja files.
#cgo LDFLAGS: -lwebrtc_magic
#cgo LDFLAGS: -lgthread-2.0 -lgtk-x11-2.0 -lgdk-x11-2.0 -lpangocairo-1.0
#cgo LDFLAGS: -latk-1.0 -lcairo -lgdk_pixbuf-2.0 -lgio-2.0 -lpangoft2-1.0 -lpango-1.0
#cgo LDFLAGS: -lgobject-2.0 -lglib-2.0 -lfontconfig -lfreetype -lX11 -lXcomposite
#cgo LDFLAGS: -lXext -lXrender -ldl -lrt -lexpat -lm
#include "cpeerconnection.h"
*/
import "C"
import (
	"fmt"
)

type Callback func()
type Obs func(Callback, Callback)

type RTCPeerConnection struct {
	// CreateOffer Obs

	// CreateAnswer func(Callback)
	// setLocalDescription
	// localDescription

	// currentLocalDescription
	// pendingLocalDescription
	// setRemoteDescription

	// remoteDescription
	// currentRemoteDescription
	// pendingRemoteDescription
	// addIceCandidate
	// signalingState
	// iceGatheringState
	// iceConnectionState
	// canTrickleIceCandidates
	// getConfiguration
	// setConfiguration
	// close
	// onnegotiationneeded
	OnIceCandidate func()
	// onsignalingstatechange
	// oniceconnectionstatechange
	// onicegatheringstatechange
	pc         C.PeerConnection
	IceServers string
}

// CreateOffer prepares ICE candidates which should be sent to the target
// peer over a signalling channel.
//
// TODO: This method blocks until success or failure occurs. Maybe it should
// be async to the user?
func (pc RTCPeerConnection) CreateOffer(success Callback, failure Callback) {
	fmt.Println("[go] creating offer...")

	// Pass return value from C through a go channel, to allow a goroutine-based
	// callback paradigm.
	// TODO(keroserene): Generalize and test this channel-based mechanism.
	r := make(chan bool, 1)
	go func() {
		success := C.CreateOffer(pc.pc)
		if 0 == success {
			r <- true
		} else {
			r <- false
		}
	}()
	status := <-r
	fmt.Println("Success: ", status)
	// Fire callbacks
	if status {
		success()
	} else {
		failure()
	}
}

// func createAnswer(pc RTCPeerConnection, c Callback) {
// C.CreateAnswer(pc.pc, c)
// }

// Install a handler for receiving ICE Candidates.
// func OnIceCandidate(pc RTCPeerConnection) {
// }

func NewPeerConnection() RTCPeerConnection {
	var ret RTCPeerConnection
	ret.pc = C.NewPeerConnection()
	// ret.IceServers = C.GetIceServers(ret.pc)
	// Assign "methods"
	// ret.CreateOffer = func(success Callback, failure Callback) {
	// createOffer(ret, success, failure)
	// }
	// ret.CreateAnswer = func(c Callback) {
	// createAnswer(ret, c)
	// }
	return ret
}

type RTCSignalingState int

/*
const {
  stable RTCSignallingState = iota
  have-local-offer
  have-remote-offer
  have-local-pranswer
  have-remote-pranswer
  closed
}
*/

type RTCConfiguration struct {
}

type RTCIceCredentialType struct {
}

type RTCIceServer struct {
	Urls       string
	Username   string
	Credential string
	// credentialType   RTCIceCredentialType
}

type RTCIceTransportPolicy struct {
}

type RTCBundlePolicy struct {
}

// func main() {
// }
