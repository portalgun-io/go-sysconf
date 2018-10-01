// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs sysconf_defs_freebsd.go

package sysconf

const (
	SC_AIO_LISTIO_MAX               = 0x2a
	SC_AIO_MAX                      = 0x2b
	SC_AIO_PRIO_DELTA_MAX           = 0x2c
	SC_ARG_MAX                      = 0x1
	SC_ATEXIT_MAX                   = 0x6b
	SC_BC_BASE_MAX                  = 0x9
	SC_BC_DIM_MAX                   = 0xa
	SC_BC_SCALE_MAX                 = 0xb
	SC_BC_STRING_MAX                = 0xc
	SC_CHILD_MAX                    = 0x2
	SC_CLK_TCK                      = 0x3
	SC_COLL_WEIGHTS_MAX             = 0xd
	SC_DELAYTIMER_MAX               = 0x2d
	SC_EXPR_NEST_MAX                = 0xe
	SC_HOST_NAME_MAX                = 0x48
	SC_IOV_MAX                      = 0x38
	SC_LINE_MAX                     = 0xf
	SC_LOGIN_NAME_MAX               = 0x49
	SC_MQ_OPEN_MAX                  = 0x2e
	SC_MQ_PRIO_MAX                  = 0x4b
	SC_NGROUPS_MAX                  = 0x4
	SC_OPEN_MAX                     = 0x5
	SC_PAGE_SIZE                    = 0x2f
	SC_PAGESIZE                     = 0x2f
	SC_RE_DUP_MAX                   = 0x10
	SC_RTSIG_MAX                    = 0x30
	SC_SEM_NSEMS_MAX                = 0x31
	SC_SEM_VALUE_MAX                = 0x32
	SC_SIGQUEUE_MAX                 = 0x33
	SC_STREAM_MAX                   = 0x1a
	SC_SYMLOOP_MAX                  = 0x78
	SC_THREAD_DESTRUCTOR_ITERATIONS = 0x55
	SC_THREAD_PRIO_INHERIT          = 0x57
	SC_THREAD_PRIO_PROTECT          = 0x58
	SC_THREAD_KEYS_MAX              = 0x56
	SC_THREAD_STACK_MIN             = 0x5d
	SC_THREAD_THREADS_MAX           = 0x5e
	SC_TIMER_MAX                    = 0x34
	SC_TTY_NAME_MAX                 = 0x65
	SC_TZNAME_MAX                   = 0x1b

	SC_ADVISORY_INFO         = 0x41
	SC_ASYNCHRONOUS_IO       = 0x1c
	SC_BARRIERS              = 0x42
	SC_CLOCK_SELECTION       = 0x43
	SC_CPUTIME               = 0x44
	SC_FSYNC                 = 0x26
	SC_IPV6                  = 0x76
	SC_JOB_CONTROL           = 0x6
	SC_MAPPED_FILES          = 0x1d
	SC_MEMLOCK               = 0x1e
	SC_MEMLOCK_RANGE         = 0x1f
	SC_MEMORY_PROTECTION     = 0x20
	SC_MESSAGE_PASSING       = 0x21
	SC_MONOTONIC_CLOCK       = 0x4a
	SC_PRIORITIZED_IO        = 0x22
	SC_PRIORITY_SCHEDULING   = 0x23
	SC_RAW_SOCKETS           = 0x77
	SC_READER_WRITER_LOCKS   = 0x4c
	SC_REALTIME_SIGNALS      = 0x24
	SC_REGEXP                = 0x4d
	SC_SAVED_IDS             = 0x7
	SC_SEMAPHORES            = 0x25
	SC_SHARED_MEMORY_OBJECTS = 0x27
	SC_SHELL                 = 0x4e
	SC_THREAD_CPUTIME        = 0x54
	SC_THREADS               = 0x60
	SC_TIMEOUTS              = 0x5f
	SC_TIMERS                = 0x29
	SC_VERSION               = 0x8

	SC_2_C_BIND    = 0x12
	SC_2_C_DEV     = 0x13
	SC_2_FORT_DEV  = 0x15
	SC_2_FORT_RUN  = 0x16
	SC_2_LOCALEDEF = 0x17
	SC_2_SW_DEV    = 0x18
	SC_2_UPE       = 0x19
	SC_2_VERSION   = 0x11

	SC_XOPEN_CRYPT       = 0x6c
	SC_XOPEN_REALTIME    = 0x6f
	SC_XOPEN_STREAMS     = 0x72
	SC_XOPEN_UNIX        = 0x73
	SC_XOPEN_VERSION     = 0x74
	SC_XOPEN_XCU_VERSION = 0x75

	SC_PHYS_PAGES       = 0x79
	SC_NPROCESSORS_CONF = 0x39
	SC_NPROCESSORS_ONLN = 0x3a
)

const (
	_BC_BASE_MAX      = 0x63
	_BC_DIM_MAX       = 0x800
	_BC_SCALE_MAX     = 0x63
	_BC_STRING_MAX    = 0x3e8
	_COLL_WEIGHTS_MAX = 0xa
	_EXPR_NEST_MAX    = 0x20
	_LINE_MAX         = 0x800
	_MQ_PRIO_MAX      = 0x40
	_RE_DUP_MAX       = 0xff
	_SEM_VALUE_MAX    = 0x7fffffff

	_CLK_TCK = 0x80

	_MAXHOSTNAMELEN = 0x100
	_MAXLOGNAME     = 0x21
	_MAXSYMLINKS    = 0x20
	_ATEXIT_SIZE    = 0x20

	_POSIX_ADVISORY_INFO         = 0x30db0
	_POSIX_ARG_MAX               = 0x1000
	_POSIX_ASYNCHRONOUS_IO       = 0x30db0
	_POSIX_BARRIERS              = 0x30db0
	_POSIX_CHILD_MAX             = 0x19
	_POSIX_CLOCK_SELECTION       = -0x1
	_POSIX_CPUTIME               = 0x30db0
	_POSIX_FSYNC                 = 0x30db0
	_POSIX_IPV6                  = 0x0
	_POSIX_JOB_CONTROL           = 0x1
	_POSIX_MAPPED_FILES          = 0x30db0
	_POSIX_MEMLOCK               = -0x1
	_POSIX_MEMLOCK_RANGE         = 0x30db0
	_POSIX_MEMORY_PROTECTION     = 0x30db0
	_POSIX_MESSAGE_PASSING       = 0x30db0
	_POSIX_MONOTONIC_CLOCK       = 0x30db0
	_POSIX_PRIORITIZED_IO        = -0x1
	_POSIX_PRIORITY_SCHEDULING   = 0x0
	_POSIX_RAW_SOCKETS           = 0x30db0
	_POSIX_READER_WRITER_LOCKS   = 0x30db0
	_POSIX_REALTIME_SIGNALS      = 0x30db0
	_POSIX_REGEXP                = 0x1
	_POSIX_SEM_VALUE_MAX         = 0x7fff
	_POSIX_SEMAPHORES            = 0x30db0
	_POSIX_SHARED_MEMORY_OBJECTS = 0x30db0
	_POSIX_SHELL                 = 0x1
	_POSIX_THREAD_PRIO_INHERIT   = 0x30db0
	_POSIX_THREAD_PRIO_PROTECT   = 0x30db0
	_POSIX_THREADS               = 0x30db0
	_POSIX_TIMEOUTS              = 0x30db0
	_POSIX_TIMERS                = 0x30db0
	_POSIX_VERSION               = 0x30db0

	_POSIX2_C_BIND    = 0x30db0
	_POSIX2_C_DEV     = -0x1
	_POSIX2_LOCALEDEF = -0x1
	_POSIX2_SW_DEV    = -0x1
	_POSIX2_UPE       = 0x30db0
	_POSIX2_VERSION   = 0x30a2c

	_XOPEN_CRYPT    = -0x1
	_XOPEN_REALTIME = -0x1
	_XOPEN_UNIX     = -0x1

	_PTHREAD_DESTRUCTOR_ITERATIONS = 0x4
	_PTHREAD_KEYS_MAX              = 0x100
	_PTHREAD_STACK_MIN             = 0x800
)

const (
	_PC_NAME_MAX = 0x4

	_PATH_DEV      = "/dev/"
	_PATH_ZONEINFO = "/usr/share/zoneinfo"
)
