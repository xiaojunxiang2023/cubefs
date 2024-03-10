
#ifndef EPOLL_H_
#define EPOLL_H_

#include<stdio.h>
#include<unistd.h>
#include<sys/types.h>
#include<sys/socket.h>
#include<netinet/in.h>
#include<arpa/inet.h>
#include<sys/epoll.h>
#include<fcntl.h>
#include<stdlib.h>

static int set_noblock(int fd) {
    int fl = fcntl(fd,F_GETFL);
    return fcntl(fd,F_SETFL,fl|O_NONBLOCK);
}

typedef void (*EventCallBack)(void *ctx);
struct EpollEvent {
    uint32_t index;
    EventCallBack cb;
    void* ctx;
};

static int epoll_fd  = -1;
static pthread_mutex_t mutex;
static pthread_t epoll_thread;
static struct EpollEvent all_event[1024] = {};
static int all_fds[1024] = {};

static void * epoll_worker(void *ctx) {
    struct epoll_event ready_ev[128];//申请空间来放就绪的事件。
    int maxnum = 128;
    int timeout = -1;//设置超时时间，若为-1，则永久阻塞等待。
    while(1) {
        int ret = epoll_wait(epoll_fd,ready_ev,maxnum,timeout);
        if(ret == -1) {
            if (errno == EINTR) {
                continue;
            }
            printf("epoll_wait fd %d ret=%d %d %s\n", epoll_fd, ret, errno, strerror(errno));
            return NULL;
        }
        for(int i = 0; i < ret; ++i) {
            uint32_t index = ready_ev[i].data.fd;
            all_event[index].cb(all_event[index].ctx);
        }
    }

    return NULL;
}

static int get_epoll_fd() {
    if (pthread_mutex_lock(&mutex) != 0){
        printf("create epoll fd %d\n", epoll_fd);
        return -1;
    }

    if (epoll_fd == -1) {
        epoll_fd = epoll_create(256);
        printf("create epoll fd %d\n", epoll_fd);
        if (pthread_create(&epoll_thread, NULL, epoll_worker, NULL) != 0){
                                printf("thread create failed!\n");
                                return -1;
        }
    }

    pthread_mutex_unlock(&mutex);
    return epoll_fd;
}

static struct EpollEvent* GetEpollEvent() {

    for(int i = 0; i < 1024; i++) {
        if (all_event[i].ctx == NULL) {
            all_event[i].index = i;
            return &all_event[i];
        }
    }

    return NULL;
}

static void DelEpollEvent(int fd) {
    for(int i = 0;i < 1024; i++) {
        if (all_fds[i] == fd) {
            all_fds[i] = 0;
            all_event[i].ctx = NULL;
            all_event[i].cb = NULL;
        }
    }
}

static int epoll_rdma_event_add(int fd, void* ctx, EventCallBack cb) {
    int epoll_fd = get_epoll_fd();
    if(epoll_fd < 0){
        return -1;
    }

    struct EpollEvent* ev = GetEpollEvent();
    ev->ctx = ctx;
    ev->cb = cb;

    all_fds[ev->index] = fd;

    struct epoll_event ep_ev;
    ep_ev.events = EPOLLIN; //| EPOLLET;
    ep_ev.data.fd = ev->index;

    if(epoll_ctl(epoll_fd, EPOLL_CTL_ADD, fd, &ep_ev) < 0){
        return -1;
    }

    return 0;
}

#endif